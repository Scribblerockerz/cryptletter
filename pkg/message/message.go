package message

import (
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
