package messages

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/allancapistrano/tangle-client-go/utils"
	iotago "github.com/iotaledger/iota.go/v2"
)

type Message struct {
	Index   string `json:"index"`
	Content string `json:"content"`
}

// Get all messages available on the node by a given index.
func GetAllMessagesByIndex(nodeUrl string, index string) ([]Message, error) {
	node := iotago.NewNodeHTTPAPIClient(nodeUrl)

	msgIdsResponse, err := node.MessageIDsByIndex(
		context.Background(),
		[]byte(index),
	)

	if err != nil {
		return nil, errors.New("unable to get message IDs")
	}

	var i uint32
	var messages []Message

	if msgIdsResponse.Count > 0 {
		for i = 0; i < msgIdsResponse.Count; i++ {
			messageId, err := iotago.MessageIDFromHexString(msgIdsResponse.MessageIDs[i])
			if err != nil {
				return nil, errors.New("unable to convert message ID from hex to message ID representation")
			}

			messageReturned, err := node.MessageByMessageID(context.Background(), messageId)
			if err != nil {
				return nil, errors.New("unable to get message by given message ID")
			}

			message, err := formatMessagePayload(*messageReturned, index)
			if err != nil {
				log.Println(err)

				message = Message{
					Index:   "Error",
					Content: err.Error(),
				}
			}

			messages = append(messages, message)
		}
	} else {
		log.Println("No messages with this index were found.")
	}

	return messages, nil
}

// Formats the message payload into a custom message type.
func formatMessagePayload(message iotago.Message, messageIndex string) (Message, error) {
	payloadInString := utils.SerializeMessagePayload(message.Payload, true)
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
			log.Panic("Unexpected array length.")
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
