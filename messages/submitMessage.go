package messages

import (
	"context"
	"log"
	"time"

	iotago "github.com/iotaledger/iota.go/v2"
	"github.com/larsid/tangle-client-go/client"
)

// Sends a new message to the Tangle Hornet Network, using a specific index.
func SubmitMessage(
	ctx context.Context,
	nodeUrl string,
	index string,
	content string,
	timeoutInSeconds int,
) bool {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeoutInSeconds)*time.Second)
	defer cancel()

	node := client.ForURL(nodeUrl)

	log.Printf("[TANGLE-CLIENT] [INFO] SubmitMessage: consultando info do nó em %s", nodeUrl)

	info, err := node.Info(ctx)
	if err != nil {
		log.Printf("[TANGLE-CLIENT] [ERROR] SubmitMessage: falha ao obter info do nó: %v", err)
		return false
	}

	log.Printf("[TANGLE-CLIENT] [INFO] SubmitMessage: nó saudável=%v network_id=%s", info.IsHealthy, info.NetworkID)

	MessagePayload := &iotago.Indexation{
		Index: []byte(index),
		Data:  []byte(content),
	}

	messageBuilder, err := iotago.NewMessageBuilder().
		Payload(MessagePayload).
		Tips(ctx, node).
		NetworkIDFromString(info.NetworkID).
		ProofOfWork(ctx, info.MinPowScore).
		Build()
	if err != nil {
		log.Printf("[TANGLE-CLIENT] [ERROR] SubmitMessage: falha ao construir mensagem index=%s: %v", index, err)
		return false
	}

	log.Printf("[TANGLE-CLIENT] [INFO] SubmitMessage: enviando mensagem index=%s payload_bytes=%d", index, len(content))

	if _, err := node.SubmitMessage(ctx, messageBuilder); err != nil {
		log.Printf("[TANGLE-CLIENT] [ERROR] SubmitMessage: falha ao submeter mensagem index=%s: %v", index, err)
		return false
	}

	log.Printf("[TANGLE-CLIENT] [INFO] SubmitMessage: mensagem submetida com sucesso index=%s", index)
	return true
}
