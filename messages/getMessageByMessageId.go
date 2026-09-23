package messages

import (
	"context"
	"errors"
	"log"

	iotago "github.com/iotaledger/iota.go/v2"
	"github.com/larsid/tangle-client-go/client"
)

// Get a message on the node by a given message ID.
func getMessageByMessageID(ctx context.Context, node *iotago.NodeHTTPAPIClient, messageIdHex string) (*iotago.Message, error) {
	messageId, err := iotago.MessageIDFromHexString(messageIdHex)
	if err != nil {
		log.Printf("[TANGLE-CLIENT] [ERROR] getMessageByMessageID: ID inválido messageId=%s: %v", messageIdHex, err)
		return &iotago.Message{}, errors.New("unable to convert message ID from hex to message ID representation")
	}

	log.Printf("[TANGLE-CLIENT] [INFO] getMessageByMessageID: buscando messageId=%s", messageIdHex)

	messageReturned, err := node.MessageByMessageID(ctx, messageId)
	if err != nil {
		log.Printf("[TANGLE-CLIENT] [ERROR] getMessageByMessageID: falha ao buscar messageId=%s: %v", messageIdHex, err)
		return &iotago.Message{}, errors.New("unable to get message by given message ID")
	}

	log.Printf("[TANGLE-CLIENT] [INFO] getMessageByMessageID: messageId=%s obtido com sucesso", messageIdHex)

	return messageReturned, nil
}

// Get a message on the node by a given message ID, into a custom message type.
func GetMessageFormattedByMessageID(ctx context.Context, nodeUrl string, messageIdHex string) (Message, error) {
	var message Message

	node := client.ForURL(nodeUrl)
	messageReturned, err := getMessageByMessageID(ctx, node, messageIdHex)
	if err != nil {
		log.Printf("[TANGLE-CLIENT] [ERROR] GetMessageFormattedByMessageID: messageId=%s: %v", messageIdHex, err)

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

	return message, nil
}
