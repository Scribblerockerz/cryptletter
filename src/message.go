package main

import (
	"fmt"
	"time"
)

// Message is a core type
type Message struct {
	Content      string
	Token        string
	CreatedAt    time.Time
	Lifetime     int64
	AccessableIP string
}

// Find something
func (m *Message) Find(token string) {
	fmt.Printf("Search for message with token: %s\n", token)
}
