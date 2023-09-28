package messages

import "github.com/allancapistrano/tangle-client-go/utils"

type Message struct {
	Index   string `json:"index"`
	Content string `json:"content"`
}

// Sanitizes a message.
func sanitizeMessage(message *Message) {
	message.Content = utils.SanitizeString(message.Content)
	message.Index = utils.SanitizeString(message.Index)
}
