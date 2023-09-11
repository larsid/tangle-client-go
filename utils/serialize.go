package utils

import (
	"fmt"
	"log"

	"github.com/iotaledger/hive.go/serializer"
	iotago "github.com/iotaledger/iota.go/v2"
)

// Serializes a given message payload using iota.go package. You can turn
// on/off the debug messages.
func SerializeMessagePayload(messagePayload serializer.Serializable, debugMode bool) string {
	messagePayloadSerialized, err := messagePayload.Serialize(serializer.DeSeriModePerformLexicalOrdering)
	if err != nil {
		log.Fatal("Unable to serialize the given message payload.")
	}

	messagePayloadInString := string(messagePayloadSerialized)

	if debugMode {
		fmt.Printf("Message serialized: %v\n", messagePayloadSerialized)
		fmt.Printf("Message in string: %s\n\n", messagePayloadInString)
	}

	return messagePayloadInString
}

// Serializes a given message using iota.go package. You can turn on/off the
// debug messages.
func SerializeMessage(message iotago.Message, debugMode bool) string {
	messageSerialized, err := message.Serialize(serializer.DeSeriModePerformLexicalOrdering)
	if err != nil {
		log.Fatal("Unable to serialize the given message.")
	}

	messageInString := string(messageSerialized)

	if debugMode {
		fmt.Printf("Message serialized: %v\n", messageSerialized)
		fmt.Printf("Message in string: %s\n\n", messageInString)
	}

	return messageInString
}
