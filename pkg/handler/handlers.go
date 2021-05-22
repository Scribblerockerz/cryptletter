package handler

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/Scribblerockerz/cryptletter/pkg/attachment"
	"github.com/Scribblerockerz/cryptletter/pkg/database"
	"github.com/Scribblerockerz/cryptletter/pkg/logger"
	"github.com/Scribblerockerz/cryptletter/pkg/message"
	"github.com/Scribblerockerz/cryptletter/pkg/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func loadMessageFromRedis(token string) (*message.Message, error) {
	hasResults, err := database.RedisClient.Exists(token).Result()
	if err != nil {
		return nil, err
	}

	// Key not found
	if hasResults == 0 {
		return nil, nil
	}

	result, err := database.RedisClient.Get(token).Result()
	if err != nil {
		return nil, err
	}

	loadedMessage := &message.Message{}
	err = json.Unmarshal([]byte(result), loadedMessage)
	if err != nil {
		return nil, err
	}

	return loadedMessage, nil
}

type responseAttachmentType struct {
	Token    string `json:"token"`
	Name     string `json:"name"`
	MimeType string `json:"mimeType"`
	Size     string `json:"size"`
}

type responseMessageType struct {
	Message              string                   `json:"message"`
	ActiveUntilTimestamp string                   `json:"activeUntilTimestamp"`
	Token                string                   `json:"token"`
	Attachments          []responseAttachmentType `json:"attachments"`
}

// ShowAction handles a single message
func ShowAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	loadedMessage, err := loadMessageFromRedis(vars["token"])
	if err != nil || loadedMessage == nil {
		NotFound(w, r)
		return
	}

	visitorHash := getHashedIP(r, loadedMessage.Token)

	// First time someone access this message, update message with new expire date
	if loadedMessage.AccessibleIP == "" {
		loadedMessage.AccessibleIP = visitorHash

		bytes, err := json.Marshal(&loadedMessage)
		if err != nil {
			panic(err)
		}

		err = database.RedisClient.Set(loadedMessage.Token, string(bytes), time.Duration(loadedMessage.Lifetime)*time.Minute).Err()

		// Determine host type from the stored attachment
		hostType := ""
		if len(loadedMessage.Attachments) > 0 {
			hostType = loadedMessage.Attachments[0].HostType
		}
		attachmentHandler := attachment.NewAttachmentHandler(hostType)

		for _, att := range loadedMessage.Attachments {
			err2 := attachmentHandler.SetTTL(att.FileID, loadedMessage.Lifetime * 60)
			if err2 != nil {
				logger.LogError(err2)
			}
		}

		if err != nil {
			fmt.Printf("An error accoured %s", err)
		}

	}

	// Is the current visitor bound with this message?
	if loadedMessage.AccessibleIP != visitorHash {
		NotFound(w, r)
		return
	}

	duration, err := database.RedisClient.TTL(loadedMessage.Token).Result()
	if err != nil {
		panic(err)
	}

	var responseAttachments []responseAttachmentType
	for _, att := range loadedMessage.Attachments {
		responseAttachments = append(responseAttachments, responseAttachmentType{
			Token:    att.Token,
			Name:     att.Name,
			MimeType: att.MimeType,
			Size:     att.Size,
		})
	}

	res, err := json.Marshal(responseMessageType{
		Message:              loadedMessage.Content,
		ActiveUntilTimestamp: strconv.FormatInt(time.Now().Add(duration).Unix()*1000, 10),
		Token:                vars["token"],
		Attachments:          responseAttachments,
	})

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(res))
}

// GetAttachmentAction handles attachment access
func GetAttachmentAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	vars := mux.Vars(r)

	loadedMessage, err := loadMessageFromRedis(vars["token"])
	if err != nil || loadedMessage == nil {
		NotFound(w, r)
		return
	}

	visitorHash := getHashedIP(r, loadedMessage.Token)

	// Is the current visitor bound with this message?
	if loadedMessage.AccessibleIP != visitorHash {
		NotFound(w, r)
		return
	}

	var data string

	for _, att := range loadedMessage.Attachments {
		if att.Token != vars["attachmentToken"] {
			continue
		}

		attachmentHandler := attachment.NewAttachmentHandler(att.HostType)
		data, err = attachmentHandler.Get(att.FileID)
		if err != nil {
			logger.LogError(err)
			NotFound(w, r)
		}
		break
	}

	fmt.Fprintf(w, data)
}

// NotFound handles a single message
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

// DeleteMessageAction will delete a message
func DeleteMessageAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	loadedMessage, err := loadMessageFromRedis(vars["token"])
	if err != nil || loadedMessage == nil {
		w.Write([]byte("{}"))
		return
	}

	visitorHash := getHashedIP(r, loadedMessage.Token)

	if loadedMessage.AccessibleIP == visitorHash {

		for _, att := range loadedMessage.Attachments {
			// TODO: determine the type of the handler based on attachment.HostType
			attachmentHandler := attachment.NewAttachmentHandler(att.HostType)
			attachmentHandler.Delete(att.FileID)
		}

		err = database.RedisClient.Del(loadedMessage.Token).Err()
		if err != nil {
			fmt.Printf("An error accoured %s", err)
		}
	}

	w.Write([]byte("{}"))
}

type requestAttachmentType struct {
	Name     string
	MimeType string
	Size     string
	Data     string
}

type requestMessageType struct {
	Delay                       int64 //`json:",string"`
	Message                     string
	CreationRestrictionPassword string //`json:",string"`
	Attachments                 []requestAttachmentType
}

// NewMessageAction will handle new messages
func NewMessageAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	decoder := json.NewDecoder(r.Body)
	requestMessage := requestMessageType{}

	err := decoder.Decode(&requestMessage)
	if err != nil {
		panic(err)
	}

	passwordProtection := viper.GetString("app.creation_protection_password")

	// Restrict letter creation with a password
	if requestMessage.CreationRestrictionPassword != passwordProtection {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// TODO: Only enable attachments if they are enabled by configuration at start
	attachmentHandler := attachment.NewAttachmentHandler(viper.GetString("app.attachments.driver"))

	var newAttachments []message.Attachment

	for _, requestAttachment := range requestMessage.Attachments {
		// Handle file storage based on current handler
		fileID, err := attachmentHandler.Put(requestAttachment.Data)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Stored new file %s at %s - %d\n", requestAttachment.Name, fileID, requestAttachment.Size)

		newAttachments = append(newAttachments, message.Attachment{
			Token:    utils.GenerateToken(),
			Name:     requestAttachment.Name,
			MimeType: requestAttachment.MimeType,
			Size:     requestAttachment.Size,
			FileID:   fileID,
		})
	}

	newMessage := message.Message{
		Content:     requestMessage.Message,
		Lifetime:    requestMessage.Delay,
		Token:       utils.GenerateToken(),
		CreatedAt:   time.Now(),
		Attachments: newAttachments,
	}

	bytes, err := json.Marshal(&newMessage)
	if err != nil {
		panic(err)
	}

	err = database.RedisClient.Set(newMessage.Token, string(bytes), time.Duration(viper.GetInt32("app.default_message_ttl"))*time.Minute).Err()
	if err != nil {
		fmt.Printf("An error accoured %s", err)
	}

	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(map[string]string{
		"token": newMessage.Token,
	})

	if err != nil {
		panic(err)
	}

	w.Write(res)
}

// GetHashedIP will return a hash of the remote ip unique to the token
func getHashedIP(req *http.Request, token string) string {

	ip := req.Header.Get("x-forwarded-for")

	if ip == "" {
		ip = strings.Split(req.RemoteAddr, ":")[0]
	}

	hash := fmt.Sprintf("%x", md5.Sum([]byte(ip+token)))
	fmt.Printf("Remote IP %s - %s - Hash: %s\n", ip, token, hash)

	return hash
}



func DecorateCORSHeadersHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
		if r.Method == http.MethodOptions {
			return
		}

		h.ServeHTTP(w, r)
	})
}
