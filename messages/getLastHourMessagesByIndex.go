package messages

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	iotago "github.com/iotaledger/iota.go/v2"
	"github.com/larsid/tangle-client-go/client"
)

// Get a limited amount of messages created in the last hour, available on the
// node by a given index.
func GetLastHourMessagesByIndex(ctx context.Context, nodeUrl string, index string, maxMessages int) ([]Message, error) {
	node := client.ForURL(nodeUrl)

	log.Printf("[TANGLE-CLIENT] [INFO] GetLastHourMessagesByIndex: buscando IDs index=%s maxMessages=%d", index, maxMessages)

	msgIdsResponse, err := node.MessageIDsByIndex(
		ctx,
		[]byte(index),
	)

	if err != nil {
		log.Printf("[TANGLE-CLIENT] [ERROR] GetLastHourMessagesByIndex: falha ao obter IDs index=%s: %v", index, err)
		return nil, errors.New("unable to get message IDs")
	}

	log.Printf("[TANGLE-CLIENT] [INFO] GetLastHourMessagesByIndex: index=%s total_ids=%d", index, msgIdsResponse.Count)

	var messages []Message

	if msgIdsResponse.Count > 0 {
		for i := uint32(0); i < msgIdsResponse.Count; i++ {
			var message Message

			messageReturned, err := getMessageByMessageID(ctx, node, msgIdsResponse.MessageIDs[i])

			if err != nil {
				log.Printf("[TANGLE-CLIENT] [ERROR] GetLastHourMessagesByIndex: falha ao buscar messageId=%s index=%s: %v", msgIdsResponse.MessageIDs[i], index, err)

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

		if len(messages) == 0 {
			log.Printf("[TANGLE-CLIENT] [INFO] GetLastHourMessagesByIndex: nenhuma mensagem na última hora index=%s", index)
		}
	} else {
		log.Printf("[TANGLE-CLIENT] [INFO] GetLastHourMessagesByIndex: nenhuma mensagem encontrada index=%s", index)
	}

	log.Printf("[TANGLE-CLIENT] [INFO] GetLastHourMessagesByIndex: index=%s mensagens_retornadas=%d", index, len(messages))
	return messages, nil
}
