package server

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
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

type jsonCommand struct {
	clientID string   `json:"clientid"`
	command  string   `json:"command"`
	args     []string `json:"arguements"`
}

func newCommand(command string) jsonCommand {
	return jsonCommand{
		command: command,
	}
}

func write(ws *websocket.Conn, clientID, cmd string, args ...string) {
	cmdStruct := newCommand(cmd)
	if args != nil {
		for _, arg := range args {
			cmdStruct.args = append(cmdStruct.args, arg)
		}
	}
	rawMessage := structToByte(cmdStruct)
	ws.WriteMessage(websocket.TextMessage, rawMessage)
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
