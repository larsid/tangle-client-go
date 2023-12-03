package messages

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	iotago "github.com/iotaledger/iota.go/v2"
)

// Get a limited amount of messages created in the last hour, available on the 
// node by a given index.
func GetLastHourMessagesByIndex(nodeUrl string, index string, maxMessages int) ([]Message, error) {
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
					Index: "Error",
					Data:  err.Error(),
				}
			} else {
				indexationPayload := messageReturned.Payload.(*iotago.Indexation)

				message = Message{
					Index: string(indexationPayload.Index),
					Data:  string(indexationPayload.Data),
				}
			}

			var data map[string]interface{}
			var createdAtInt64 int64

			err = json.Unmarshal([]byte(message.Data), &data)

			if err != nil {
				return nil, errors.New("error trying to decode JSON")
			}

			if createdAt, ok := data["createdAt"].(float64); ok {
				// One-hour time limit
				timeLimit := time.Now().UnixMilli() - 1*60*60*1000

				createdAtInt64 = int64(createdAt)

				if createdAtInt64 >= timeLimit {
					messages = append(messages, message)
				}

				if len(messages) == maxMessages {
					break
				}
			} else {
				return nil, errors.New("error, this JSON doesn't have 'createdAt' parameter")
			}
		}

		if (len(messages) == 0) {
			log.Println("No messages have been created in the last hour.")
		}
	} else {
		log.Println("No messages with this index were found.")
	}

	return messages, nil
}
