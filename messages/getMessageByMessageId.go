package messages

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/allancapistrano/tangle-client-go/utils"
	iotago "github.com/iotaledger/iota.go/v2"
)

// Get a message on the node by a given message ID.
func getMessageByMessageID(nodeUrl string, messageIdHex string) (*iotago.Message, error) {
	node := iotago.NewNodeHTTPAPIClient(nodeUrl)

	messageId, err := iotago.MessageIDFromHexString(messageIdHex)
	if err != nil {
		return &iotago.Message{}, errors.New("unable to convert message ID from hex to message ID representation")
	}

	messageReturned, err := node.MessageByMessageID(context.Background(), messageId)
	if err != nil {
		return &iotago.Message{}, errors.New("unable to get message by given message ID")
	}

	return messageReturned, nil
}

// Get a message on the node by a given message ID, into a custom message type.
func GetMessageFormattedByMessageID(nodeUrl string, messageIdHex string) (Message, error) {
	var message Message

	messageReturned, err := getMessageByMessageID(nodeUrl, messageIdHex)
	if err != nil {
		log.Println(err)

		message = Message{
			Index:   "Error",
			Content: err.Error(),
		}
	} else {
		message, err = formatMessagePayloadWithoutIndex(messageReturned)
		if err != nil {
			log.Println(err)

			message = Message{
				Index:   "Error",
				Content: err.Error(),
			}
		}

		SanitizeMessage(&message)
	}

	return message, nil
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
