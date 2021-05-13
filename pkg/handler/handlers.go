package handler

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Scribblerockerz/cryptletter/pkg/database"
	"github.com/Scribblerockerz/cryptletter/pkg/message"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// IndexAction handles the homepage
func IndexAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serve static public order")
}

type responseAttachmentType struct {
	Name     string `json:"name"`
	MimeType string `json:"mimeType"`
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

	hasResults, err1 := database.RedisClient.Exists(vars["token"]).Result()
	if err1 != nil {
		panic(err1)
	}

	// Key not found
	if hasResults == 0 {
		NotFound(w, r)
		return
	}

	result, err := database.RedisClient.Get(vars["token"]).Result()
	if err != nil {
		panic(err)
	}

	loadedMessage := &message.Message{}
	err = json.Unmarshal([]byte(result), loadedMessage)
	if err != nil {
		panic(err)
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
	for _, attachment := range loadedMessage.Attachments {
		responseAttachments = append(responseAttachments, responseAttachmentType{
			Name:     attachment.Name,
			MimeType: attachment.MimeType,
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

// NotFound handles a single message
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

// DeleteMessageAction will delete a message
func DeleteMessageAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	hasResults, err1 := database.RedisClient.Exists(vars["token"]).Result()
	if err1 != nil {
		panic(err1)
	}

	// Key not found
	if hasResults == 0 {
		w.Write([]byte("{}"))
		return
	}

	result, err := database.RedisClient.Get(vars["token"]).Result()
	if err != nil {
		panic(err)
	}

	loadedMessage := &message.Message{}
	err = json.Unmarshal([]byte(result), loadedMessage)
	if err != nil {
		panic(err)
	}

	visitorHash := getHashedIP(r, loadedMessage.Token)

	// First time somone access this message, update message with new expire date

	if loadedMessage.AccessibleIP == visitorHash {

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

	var newAttachments []message.Attachment

	for _, requestAttachment := range requestMessage.Attachments {
		newAttachments = append(newAttachments, message.Attachment{
			Name:     requestAttachment.Name,
			MimeType: requestAttachment.MimeType,
		})
	}

	newMessage := message.Message{
		Content:     requestMessage.Message,
		Lifetime:    requestMessage.Delay,
		Token:       generateToken(),
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

func generateToken() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", sha256.Sum256([]byte(base64.StdEncoding.EncodeToString(randomBytes)[:32])))
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
