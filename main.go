package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/iotaledger/hive.go/serializer"
	iotago "github.com/iotaledger/iota.go/v2"
	// "time"
)

func main() {
	nodeURL := "http://127.0.0.1:14265"

	// Informações da Rede

	node := iotago.NewNodeHTTPAPIClient(nodeURL)

	info, err := node.Info(context.Background())

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(info.Name)
	fmt.Println(info.NetworkID)
	fmt.Println(info.MinPowScore)

	// Transações

	// Publicando uma transação

	// payload := &iotago.Indexation{
	// 	Index: []byte("Hello_World"),
	// 	// Data:  []byte("{\"ID\": \"Device Test\", \"published_at\" 123456789}"),
	// 	Data:  []byte("ahgshjsduyr"),
	// }

	// ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	// defer cancelFunc()

	// msg, err := iotago.NewMessageBuilder().
	// 	Payload(payload).
	// 	Tips(context.Background(), node).
	// 	NetworkIDFromString(info.NetworkID).
	// 	ProofOfWork(ctx, info.MinPowScore).
	// 	Build()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// if _, err := node.SubmitMessage(context.Background(), msg); err != nil {
	// 	log.Panic(err)
	// }

	// Lendo a transação pelo Index

	msgReturnedId, err := node.MessageIDsByIndex(
		context.Background(), 
		[]byte("Hello World"),
	)

	if err != nil {
		log.Panic(err)
	}

	msgId, err := iotago.MessageIDFromHexString(msgReturnedId.MessageIDs[4])
	if err != nil {
		log.Panic(err)
	}

	msgReturned, err := node.MessageByMessageID(context.Background(), msgId)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Qtd mensagens: %d\n", msgReturnedId.Count)
	fmt.Printf("Index: %s\n", msgReturnedId.Index)
	fmt.Printf("IDs das mensagens: %s\n", msgReturnedId.MessageIDs)

	payloadSerialized, err := msgReturned.Payload.Serialize(serializer.DeSeriModePerformLexicalOrdering)
	if err != nil {
		log.Panic(err)
	}

	// fmt.Println(msgReturned.Payload)

	fmt.Println(payloadSerialized)

	payloadString := string(payloadSerialized)

	fmt.Println(payloadString)

	if (strings.Contains(payloadString, "/")) {
		payloadTemp := strings.Split(payloadString, "/")

		fmt.Println(payloadTemp[0])
		fmt.Println(payloadTemp[1])
	} else if (strings.Contains(payloadString, "\v")) {
		payloadTemp := strings.Split(payloadString, "\v")

		if (len(payloadTemp) == 2) {
			fmt.Println(payloadTemp[0])
			fmt.Println(payloadTemp[1])
		} else if (len(payloadTemp) == 3) {
			fmt.Println(payloadTemp[1])
			fmt.Println(payloadTemp[2])
		} else {
			log.Panic("Unexpected array length.")
		}
	} else {
		fmt.Println(payloadString)

		log.Fatal("Malformed payload.")
	}
}
