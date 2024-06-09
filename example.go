package main

import (
	"fmt"
	"log"

	infoNode "github.com/larsid/tangle-client-go/info"
	"github.com/larsid/tangle-client-go/messages"
)

func main() {
	nodeURL := "http://127.0.0.1:14265"

	// Network info
	nodeInfo, err := infoNode.GetNodeInfo(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(nodeInfo)

	// All Network info
	allNodeInfo, err := infoNode.GetAllNodeInfo(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(allNodeInfo)

	// Submitting some message.
	messages.SubmitMessage(nodeURL, "LB_STATUS", "{\"available\":true,\"avgLoad\":3,\"createdAt\":1695652263921,\"group\":\"group3\",\"lastLoad\":4,\"publishedAt\":1695652267529,\"source\":\"source4\",\"type\":\"LB_STATUS\"}", 15)

	// Reading some messages by an index.
	messagesByIndex, err := messages.GetAllMessagesByIndex(nodeURL, "LB_STATUS")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range messagesByIndex {
		fmt.Printf("Index: %s | Data: %s\n", v.Index, v.Data)
	}

	messageID := "d57c9ad40b7079fd8e36cd3d127b3aed9fff7e3f293f1fe1913b4d850ba0814d"

	message, err := messages.GetMessageFormattedByMessageID(
		nodeURL,
		messageID,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message.Index)
	fmt.Println(message.Data)

	// Reading max of three messages by an index.
	limitedMessages, err := messages.GetLastHourMessagesByIndex(nodeURL, "92015a2d-4bae-428d-a428-6338f465e72c", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(limitedMessages)
}
