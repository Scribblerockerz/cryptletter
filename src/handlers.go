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

	fmt.Fprintf(w, RenderLayout("show.hbs", map[string]string{
		"message":              loadedMessage.Content,
		"activeUntilTimestamp": strconv.FormatInt(time.Now().Local().Add(time.Minute*time.Duration(15)).Unix()*1000, 10),
		"token":                vars["token"],
	}))

	// val, err := RedisClient.Get(vars["token"]).Result()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf(val)
	// return

	// loadedMessage := &Message{}
	// err = json.Unmarshal([]byte(val), loadedMessage)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(vars["token"], time.Now().Local().Add(time.Second*time.Duration(15)).Unix())

	// if false {
	// 	NotFound(w, r)
	// 	return
	// }

	// result := RenderLayout("show.hbs", map[string]string{
	// 	"message":              "Hello World",
	// 	"activeUntilTimestamp": strconv.FormatInt(time.Now().Local().Add(time.Minute*time.Duration(15)).Unix()*1000, 10),
	// 	"token":                vars["token"],
	// })
	// fmt.Fprintf(w, result)
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

// NewMessageAction will handle new messages
func NewMessageAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	delay, _ := strconv.ParseInt(r.FormValue("delay"), 10, 64)

	message := Message{
		Content:   r.FormValue("message"),
		Lifetime:  delay * 60,
		Token:     generateToken(),
		CreatedAt: time.Now(),
	}

	fmt.Println(message)

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
