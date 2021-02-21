package messages

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
)

// ControlMessageBytesToMessage unwraps the ControlMessage bytes into its message
func ControlMessageBytesToMessage(controlMessageBytes []byte) (message interface{}, err error) {
	controlMessage, err := ControlMessageBytesToControlMessage(controlMessageBytes)
	if err != nil {
		return
	}
	message, err = getEmptyMessageObject(controlMessage.MessageType)
	if err != nil {
		log.Printf("Could not getEmptyMessageObject: %v\n", err)
		return
	}

	err = json.Unmarshal([]byte(controlMessage.Message), &message)
	if err != nil {
		log.Printf("Could not unmarshal message: %v\n", err)
		return
	}

	return
}
func ControlMessageToMessage(controlMessage ControlMessage) (message interface{}, err error) {
	message, err = getEmptyMessageObject(controlMessage.MessageType)
	if err != nil {
		log.Printf("Could not getEmptyMessageObject: %v\n", err)
		return
	}

	err = json.Unmarshal([]byte(controlMessage.Message), &message)
	if err != nil {
		log.Printf("Could not unmarshal message: %v\n", err)
		return
	}
	return
}
func ControlMessageBytesToControlMessage(controlMessageBytes []byte) (controlMessage ControlMessage, err error) {
	err = json.Unmarshal(controlMessageBytes, &controlMessage)
	if err != nil {
		log.Printf("Could not unmarshal message: %v\n", err)
		return
	}
	return
}

func ControlMessageToControlMessageBytes(controlMessage interface{}) (controlMessageBytes []byte, err error) {
	controlMessageBytes, err = json.Marshal(controlMessage)
	if err != nil {
		log.Printf("Could not marshal control message: %v\n", err)
		return
	}
	return
}

// MessageToControlMessageBytes wraps any message with the bytes representation of a ControlMessage
func MessageToControlMessageBytes(message interface{}) (controlMessageBytes []byte, err error) {
	messageJSONBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Could not marshal message: %v\n", err)
		return
	}
	controlMessage := ControlMessage{
		MessageType: reflect.TypeOf(message).String(),
		Message:     string(messageJSONBytes),
	}
	controlMessageBytes, err = json.Marshal(controlMessage)
	if err != nil {
		log.Printf("Could not marshal control message: %v\n", err)
		return
	}
	return
}

// getEmptyMessageObject returns an empty struct of type controlType
func getEmptyMessageObject(controlType string) (interface{}, error) {
	messageType, ok := controlMessageTypeRegister[controlType]
	if !ok {
		return nil, errors.New("Control message object not registered")
	}
	return reflect.New(messageType).Elem().Interface(), nil
}

// func newMessage(controlMessage ControlMessage) {

// }
