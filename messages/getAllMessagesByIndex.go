package messages

import (
	"context"
	"errors"
	"log"

	iotago "github.com/iotaledger/iota.go/v2"
	"github.com/larsid/tangle-client-go/client"
)

// Get all messages available on the node by a given index.
func GetAllMessagesByIndex(ctx context.Context, nodeUrl string, index string) ([]Message, error) {
	node := client.ForURL(nodeUrl)

	log.Printf("[TANGLE-CLIENT] [INFO] GetAllMessagesByIndex: buscando IDs index=%s", index)

	msgIdsResponse, err := node.MessageIDsByIndex(
		ctx,
		[]byte(index),
	)

	if err != nil {
		log.Printf("[TANGLE-CLIENT] [ERROR] GetAllMessagesByIndex: falha ao obter IDs index=%s: %v", index, err)
		return nil, errors.New("unable to get message IDs")
	}

	log.Printf("[TANGLE-CLIENT] [INFO] GetAllMessagesByIndex: index=%s total_ids=%d", index, msgIdsResponse.Count)

	var messages []Message

	if msgIdsResponse.Count > 0 {
		for i := uint32(0); i < msgIdsResponse.Count; i++ {
			var message Message

			messageReturned, err := getMessageByMessageID(ctx, node, msgIdsResponse.MessageIDs[i])

			if err != nil {
				log.Printf("[TANGLE-CLIENT] [ERROR] GetAllMessagesByIndex: falha ao buscar messageId=%s index=%s: %v", msgIdsResponse.MessageIDs[i], index, err)

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

			messages = append(messages, message)
		}
	} else {
		log.Printf("[TANGLE-CLIENT] [INFO] GetAllMessagesByIndex: nenhuma mensagem encontrada index=%s", index)
	}

	log.Printf("[TANGLE-CLIENT] [INFO] GetAllMessagesByIndex: index=%s mensagens_retornadas=%d", index, len(messages))
	return messages, nil
}
