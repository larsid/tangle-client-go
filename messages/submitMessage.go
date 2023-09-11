package messages

import (
	"context"
	"log"
	"time"

	iotago "github.com/iotaledger/iota.go/v2"
)

// Sends a new message to the Tangle Hornet Network, using a specific index. 
func SubmitMessage(
	nodeUrl string, 
	index string, 
	content string, 
	timeoutInSeconds int,
) bool {
	node := iotago.NewNodeHTTPAPIClient(nodeUrl)

	info, err := node.Info(context.Background())
	if err != nil {
		log.Println("Unable to get node information.")
		return false
	}

	MessagePayload := &iotago.Indexation {
		Index: []byte(index),
		Data:  []byte(content),
	}

	ctx, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Duration(timeoutInSeconds)*time.Second,
	)
	defer cancelFunc()

	messageBuilder, err := iotago.NewMessageBuilder().
		Payload(MessagePayload).
		Tips(context.Background(), node).
		NetworkIDFromString(info.NetworkID).
		ProofOfWork(ctx, info.MinPowScore).
		Build()
	if err != nil {
		log.Printf("Unable to create a new message builder.")
		return false
	}

	if _, err := node.SubmitMessage(context.Background(), messageBuilder); err != nil {
		log.Println("Unable to submit new message.")
		return false
	}

	return true
}
