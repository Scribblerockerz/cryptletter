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
	AccessibleIP string
	Attachments  []Attachment
}

// Attachment is a core type
type Attachment struct {
	Token    string
	Name     string
	MimeType string
	Size     string
	FileID   string
	HostType string
}
