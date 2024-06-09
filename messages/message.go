package messages

import (
	"errors"
	"strings"

	iotago "github.com/iotaledger/iota.go/v2"

	"github.com/larsid/tangle-client-go/utils"
)

type Message struct {
	Index string `json:"index"`
	Data  string `json:"data"`
}

// Sanitizes a message.
//
// Deprecated: sanitizeMessage exists for historical compatibility
// and should not be used. You no longer need to sanitize a message.
func sanitizeMessage(message *Message) {
	message.Data = utils.SanitizeString(message.Data)
	message.Index = utils.SanitizeString(message.Index)
}

// Formats the message payload into a custom message type.
//
// Deprecated: formatMessagePayload exists for historical compatibility
// and should not be used. You no longer need to format a message payload
// because now Indexation payload is used.
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
	} else if strings.Contains(payloadInString, "|") {
		payloadTemp := strings.Split(payloadInString, "|")

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
		Index: strings.Trim(index, "\f"),
		Data:  strings.Trim(content, "\f"),
	}

	return formattedMessage, nil
}

// Formats the message payload into a custom message type.
//
// Deprecated: formatMessagePayloadWithoutIndex exists for historical
// compatibility and should not be used. You no longer need to format a message
// payload because now Indexation payload is used.
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
		Index: strings.Trim(index, "\f"),
		Data:  strings.Trim(content, "\f"),
	}

	return formattedMessage, nil
}
