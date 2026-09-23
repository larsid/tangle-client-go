package messages

import (
	"context"
	"errors"
	"log"

	iotago "github.com/iotaledger/iota.go/v2"
	"github.com/larsid/tangle-client-go/client"
)

// GetMessagesByIndexWithLimit returns the most recent messages for a given index,
// fetching at most limit messages from the node.
func GetMessagesByIndexWithLimit(ctx context.Context, nodeUrl string, index string, limit int) ([]Message, error) {
	if limit <= 0 {
		return nil, errors.New("limit must be a positive integer")
	}

	node := client.ForURL(nodeUrl)

	log.Printf("[TANGLE-CLIENT] [INFO] GetMessagesByIndexWithLimit: buscando IDs index=%s limit=%d", index, limit)

	msgIdsResponse, err := node.MessageIDsByIndex(
		ctx,
		[]byte(index),
	)
	if err != nil {
		log.Printf("[TANGLE-CLIENT] [ERROR] GetMessagesByIndexWithLimit: falha ao obter IDs index=%s: %v", index, err)
		return nil, errors.New("unable to get message IDs")
	}

	log.Printf("[TANGLE-CLIENT] [INFO] GetMessagesByIndexWithLimit: index=%s total_ids=%d", index, msgIdsResponse.Count)

	if msgIdsResponse.Count == 0 {
		log.Printf("[TANGLE-CLIENT] [INFO] GetMessagesByIndexWithLimit: nenhuma mensagem encontrada index=%s", index)
		return []Message{}, nil
	}

	start := int(msgIdsResponse.Count) - limit
	if start < 0 {
		start = 0
	}
	idsToFetch := msgIdsResponse.MessageIDs[start:]

	messages := make([]Message, 0, len(idsToFetch))

	for _, messageId := range idsToFetch {
		messageReturned, err := getMessageByMessageID(ctx, node, messageId)
		if err != nil {
			log.Printf("[TANGLE-CLIENT] [ERROR] GetMessagesByIndexWithLimit: falha ao buscar messageId=%s index=%s: %v", messageId, index, err)
			messages = append(messages, Message{
				Index: "Error",
				Data:  err.Error(),
			})
			continue
		}

		indexationPayload := messageReturned.Payload.(*iotago.Indexation)
		messages = append(messages, Message{
			Index: string(indexationPayload.Index),
			Data:  string(indexationPayload.Data),
		})
	}

	log.Printf("[TANGLE-CLIENT] [INFO] GetMessagesByIndexWithLimit: index=%s mensagens_retornadas=%d", index, len(messages))
	return messages, nil
}
