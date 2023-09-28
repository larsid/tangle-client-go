package messages

import (
	"context"
	"errors"
	"log"

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

		sanitizeMessage(&message)
	}

	return message, nil
}
