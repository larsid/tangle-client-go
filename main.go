package main

import (
	"fmt"

	infoNode "github.com/allancapistrano/tangle-client-go/info"
	"github.com/allancapistrano/tangle-client-go/messages"
)

func main() {
	nodeURL := "http://127.0.0.1:14265"

	// Informações da Rede
	nodeInfo := infoNode.GetNodeInfo(nodeURL)

	fmt.Println(nodeInfo)

	// Transações

	// Publicando uma transação
	// messages.SubmitMessage(nodeURL, "LB_REPLY", "{asdfghjkl}", 15)

	// Lendo a transação pelo Index
	messages := messages.GetAllMessagesByIndex(nodeURL, "LB_REPLY")

	for _, v := range messages {
		fmt.Printf("Index: %s | Content: %s\n", v.Index, v.Content)
	}
}
