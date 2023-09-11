package messages

import (
	"context"
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
func GetAllMessagesByIndex(nodeUrl string, index string) []Message {
	node := iotago.NewNodeHTTPAPIClient(nodeUrl)

	msgIdsResponse, err := node.MessageIDsByIndex(
		context.Background(),
		[]byte(index),
	)

	if err != nil {
		log.Fatal("Unable to get message IDs.")
	}

	var i uint32
	var messages []Message

	if msgIdsResponse.Count > 0 {
		for i = 0; i < msgIdsResponse.Count; i++ {
			messageId, err := iotago.MessageIDFromHexString(msgIdsResponse.MessageIDs[i])
			if err != nil {
				log.Fatal(err)
			}

			messageReturned, err := node.MessageByMessageID(context.Background(), messageId)
			if err != nil {
				log.Fatal(err)
			}

			message := formatMessagePayload(*messageReturned)

			messages = append(messages, message)
		}
	} else {
		log.Println("No messages with this index were found.")
	}

	return messages
}

// Formats the message payload into a custom message type.
func formatMessagePayload(message iotago.Message) Message {
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
	} else {
		log.Fatal("Malformed payload.")
	}

	return Message{
		Index:   index,
		Content: content,
	}
}
