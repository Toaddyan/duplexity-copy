package main

import (
	"encoding/json"
	"log"
)

func jsonErrorHandler() {
	log.Panic("Error Marshal/Unmarshal Object")
}
func newCommand(command string) jsonCommand {
	return jsonCommand{
		ClientID: config.ClientID,
		Command:  command,
	}
}

type jsonCommand struct {
	ClientID string
	Command  string
	Args     []string
}

// byteToStruct converts rawMessage into struct
func byteToStruct(rawMessage []byte) jsonCommand {
	// For now this only converts jsonCommand
	cmdStruct := jsonCommand{}
	err := json.Unmarshal(rawMessage, &cmdStruct)
	if err != nil {
		log.Panicf("Could not unmarshal: %v", err)
	}
	return cmdStruct
}

// converts struct into []byte
func structToByte(cmd jsonCommand) []byte {
	log.Println("Converting cmd ", cmd, " to bytes")
	rawMessage, err := json.Marshal(cmd)
	if err != nil {
		log.Panicf("Could not marshal: %v", err)
	}
	log.Println("CONVERTED to string", string(rawMessage))
	return rawMessage

}
