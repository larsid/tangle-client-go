package main

import (
	"fmt"
	"log"
	"strings"

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
	// messages.SubmitMessage(nodeURL, "LB_REPLY", "{asdfghjkl}", 15) // TODO: Conseguir ler esse tipo de mensagem

	// Reading some messages by an index.
	messages, err := messages.GetAllMessagesByIndex(nodeURL, "LB_ENTRY_REPLY")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range messages {
		fmt.Println([]byte(strings.Trim(v.Content, "\t")))
		fmt.Printf("Index: %s | Content: %s\n", v.Index, v.Content)
	}
}
