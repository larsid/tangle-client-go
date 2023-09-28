package messages

import (
	"errors"
	"strings"

	iotago "github.com/iotaledger/iota.go/v2"

	"github.com/allancapistrano/tangle-client-go/utils"
)

type Message struct {
	Index   string `json:"index"`
	Content string `json:"content"`
}

// Sanitizes a message.
func sanitizeMessage(message *Message) {
	message.Content = utils.SanitizeString(message.Content)
	message.Index = utils.SanitizeString(message.Index)
}

// Formats the message payload into a custom message type.
func formatMessagePayload(message iotago.Message, messageIndex string) (Message, error) {
	payloadInString, err := utils.SerializeMessagePayload(message.Payload, true)
	if err != nil {
		return Message{}, err
	}

	index := ""
	content := ""

	if strings.Contains(payloadInString, "/") {
		payloadTemp := strings.Split(payloadInString, "/")

		index = payloadTemp[0]
		content = payloadTemp[1]
	} else if strings.Contains(payloadInString, "\v") {
		payloadTemp := strings.Split(payloadInString, "\v")

		if len(payloadTemp) == 2 {
			index = payloadTemp[0]
			content = payloadTemp[1]
		} else if len(payloadTemp) == 3 {
			index = payloadTemp[1]
			content = payloadTemp[2]
		} else {
			return Message{}, errors.New("unexpected array length")
		}
	} else if strings.Contains(payloadInString, "\t") {
		payloadTemp := strings.Split(payloadInString, "\t")

		index = payloadTemp[0]
		content = payloadTemp[1]
	} else if strings.Contains(payloadInString, messageIndex) {
		payloadTemp := strings.Split(payloadInString, messageIndex)

		index = messageIndex
		content = payloadTemp[1]
	} else {
		return Message{}, errors.New("malformed payload")
	}

	formattedMessage := Message{
		Index:   strings.Trim(index, "\f"),
		Content: strings.Trim(content, "\f"),
	}

	return formattedMessage, nil
}

// Formats the message payload into a custom message type.
func formatMessagePayloadWithoutIndex(message *iotago.Message) (Message, error) {
	payloadInString, err := utils.SerializeMessagePayload(message.Payload, true)
	if err != nil {
		return Message{}, err
	}

	index := ""
	content := ""

	if strings.Contains(payloadInString, "/") {
		payloadTemp := strings.Split(payloadInString, "/")

		index = payloadTemp[0]
		content = payloadTemp[1]
	} else if strings.Contains(payloadInString, "\v") {
		payloadTemp := strings.Split(payloadInString, "\v")

		if len(payloadTemp) == 2 {
			index = payloadTemp[0]
			content = payloadTemp[1]
		} else if len(payloadTemp) == 3 {
			index = payloadTemp[1]
			content = payloadTemp[2]
		} else {
			return Message{}, errors.New("unexpected array length")
		}
	} else if strings.Contains(payloadInString, "\t") {
		payloadTemp := strings.Split(payloadInString, "\t")

		index = payloadTemp[0]
		content = payloadTemp[1]
	} else {
		return Message{}, errors.New("malformed payload")
	}

	formattedMessage := Message{
		Index:   strings.Trim(index, "\f"),
		Content: strings.Trim(content, "\f"),
	}

	return formattedMessage, nil
}
