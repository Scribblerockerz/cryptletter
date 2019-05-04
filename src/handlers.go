package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// IndexAction handles the homepage
func IndexAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, RenderLayout("index.hbs", map[string]string{}))
}

// ShowAction handles a single message
func ShowAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	hasResults, err1 := RedisClient.Exists(vars["token"]).Result()
	if err1 != nil {
		panic(err1)
	}

	// Key not found
	if hasResults == 0 {
		NotFound(w, r)
		return
	}

	result, err := RedisClient.Get(vars["token"]).Result()
	if err != nil {
		panic(err)
	}

	loadedMessage := &Message{}
	err = json.Unmarshal([]byte(result), loadedMessage)
	if err != nil {
		panic(err)
	}

	visitorHash := getHashedIP(r, loadedMessage.Token)

	// First time somone access this message, update message with new expire date
	if loadedMessage.AccessableIP == "" {
		loadedMessage.AccessableIP = visitorHash

		bytes, err := json.Marshal(&loadedMessage)
		if err != nil {
			panic(err)
		}

		err = RedisClient.Set(loadedMessage.Token, string(bytes), time.Duration(loadedMessage.Lifetime)*time.Minute).Err()
		if err != nil {
			fmt.Printf("An error accoured %s", err)
		}

	}

	// Is the current visitor bound with this message?
	if loadedMessage.AccessableIP != visitorHash {
		NotFound(w, r)
		return
	}

	duration, err := RedisClient.TTL(loadedMessage.Token).Result()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, RenderLayout("show.hbs", map[string]string{
		"message":              loadedMessage.Content,
		"activeUntilTimestamp": strconv.FormatInt(time.Now().Add(duration).Unix()*1000, 10),
		"token":                vars["token"],
		"durationString":       duration.String(),
	}))
}

// NotFound handles a single message
func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, RenderLayout("404.hbs", map[string]string{}))
}

// StyleguideAction displayes all used elements
func StyleguideAction(w http.ResponseWriter, r *http.Request) {
	result := RenderLayout("styleguide.hbs", map[string]string{})
	fmt.Fprintf(w, result)
}

// DeleteMessageAction will delete a message
func DeleteMessageAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	hasResults, err1 := RedisClient.Exists(vars["token"]).Result()
	if err1 != nil {
		panic(err1)
	}

	// Key not found
	if hasResults == 0 {
		w.Write([]byte("{}"))
		return
	}

	result, err := RedisClient.Get(vars["token"]).Result()
	if err != nil {
		panic(err)
	}

	loadedMessage := &Message{}
	err = json.Unmarshal([]byte(result), loadedMessage)
	if err != nil {
		panic(err)
	}

	visitorHash := getHashedIP(r, loadedMessage.Token)

	// First time somone access this message, update message with new expire date

	if loadedMessage.AccessableIP == visitorHash {

		err = RedisClient.Del(loadedMessage.Token).Err()
		if err != nil {
			fmt.Printf("An error accoured %s", err)
		}
	}

	w.Write([]byte("{}"))
}

type requestMessageType struct {
	Delay   int64 `json:",string"`
	Message string
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

	message := Message{
		Content:   requestMessage.Message,
		Lifetime:  requestMessage.Delay,
		Token:     generateToken(),
		CreatedAt: time.Now(),
	}

	bytes, err := json.Marshal(&message)
	if err != nil {
		panic(err)
	}

	err = RedisClient.Set(message.Token, string(bytes), time.Duration(Config.App.DefaultTTLForNewMessages)*time.Minute).Err()
	if err != nil {
		fmt.Printf("An error accoured %s", err)
	}

	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(map[string]string{
		"token": message.Token,
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
