package main

import (
	"fmt"
	"log"

	infoNode "github.com/allancapistrano/tangle-client-go/info"
	"github.com/allancapistrano/tangle-client-go/messages"
)

func main() {
	nodeURL := "http://127.0.0.1:14265"

	// Network info
	nodeInfo, err := infoNode.GetNodeInfo(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(nodeInfo)

	// Submitting some message.
	// messages.SubmitMessage(nodeURL, "LB_REPLY", "{asdfghjkl}", 15)

	// Reading some messages by an index.
	messagesByIndex, err := messages.GetAllMessagesByIndex(nodeURL, "REP_EVALUATION")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range messagesByIndex {
		fmt.Printf("Index: %s | Content: %s\n", v.Index, v.Content)
	}

	messageID := "9597556533c5e91112c0b02244799a4d308ca007486e8e844d5d78b0f298b667"

	message, err := messages.GetMessageFormattedByMessageID(
		nodeURL,
		messageID,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message.Index)
	fmt.Println(message.Content)
}
