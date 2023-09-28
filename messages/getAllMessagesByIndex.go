package messages

import (
	"context"
	"errors"
	"log"

	iotago "github.com/iotaledger/iota.go/v2"
)

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
			var message Message

			messageReturned, err := getMessageByMessageID(nodeUrl, msgIdsResponse.MessageIDs[i])

			if err != nil {
				log.Println(err)

				message = Message{
					Index:   "Error",
					Content: err.Error(),
				}
			} else {
				message, err = formatMessagePayload(*messageReturned, index)

				if err != nil {
					log.Println(err)

					message = Message{
						Index:   "Error",
						Content: err.Error(),
					}
				} else {
					sanitizeMessage(&message)

					messages = append(messages, message)
				}
			}
		}
	} else {
		log.Println("No messages with this index were found.")
	}

	return messages, nil
}
