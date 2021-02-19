package main

import (
	"encoding/json"
	"log"
)

// Command constants
const (
	registerConnection = "registerconnection"
	getDataURI         = "getDataURI"
	updateTTL          = "updatettl"
	disconnect         = "disconnect"
)

func jsonErrorHandler() {
	log.Panic("Error Marshal/Unmarshal Object")
}
func newCommand(command string) jsonCommand {
	return jsonCommand{
		clientID: config.ClientID,
		command:  command,
	}
}

type jsonCommand struct {
	clientID string   `json:"clientid"`
	command  string   `json:"command"`
	args     []string `json:"arguements"`
}

// byteToStruct converts rawMessage into struct
func byteToStruct(rawMessage []byte) jsonCommand {
	// For now this only converts jsonCommand
	cmdStruct := jsonCommand{}
	err := json.Unmarshal(rawMessage, &cmdStruct)
	if err != nil {
		jsonErrorHandler()
	}
	return cmdStruct
}

// converts struct into []byte
func structToByte(cmd interface{}) []byte {
	rawMessage, err := json.Marshal(cmd)
	if err != nil {
		jsonErrorHandler()
	}
	return rawMessage

}
