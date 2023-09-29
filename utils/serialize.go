package utils

import (
	"errors"
	"fmt"

	"github.com/iotaledger/hive.go/serializer"
	iotago "github.com/iotaledger/iota.go/v2"
)

// Serializes a given message payload using iota.go package. You can turn
// on/off the debug messages.
//
// Deprecated: SerializeMessagePayload exists for historical compatibility
// and should not be used. You no longer need to serialize a message payload
// because now Indexation payload is used.
func SerializeMessagePayload(messagePayload serializer.Serializable, debugMode bool) (string, error) {
	messagePayloadSerialized, err := messagePayload.Serialize(serializer.DeSeriModePerformLexicalOrdering)
	if err != nil {
		return "", errors.New("unable to serialize the given message payload")
	}

	messagePayloadInString := string(messagePayloadSerialized)

	if debugMode {
		fmt.Printf("Message serialized: %v\n", messagePayloadSerialized)
		fmt.Printf("Message in string: %s\n\n", messagePayloadInString)
	}

	return messagePayloadInString, nil
}

// Serializes a given message using iota.go package. You can turn on/off the
// debug messages.
//
// Deprecated: SerializeMessage exists for historical compatibility
// and should not be used. You no longer need to serialize a message payload
// because now Indexation payload is used.
func SerializeMessage(message iotago.Message, debugMode bool) (string, error) {
	messageSerialized, err := message.Serialize(serializer.DeSeriModePerformLexicalOrdering)
	if err != nil {
		return "", errors.New("unable to serialize the given message")
	}

	messageInString := string(messageSerialized)

	if debugMode {
		fmt.Printf("Message serialized: %v\n", messageSerialized)
		fmt.Printf("Message in string: %s\n\n", messageInString)
	}

	return messageInString, nil
}
