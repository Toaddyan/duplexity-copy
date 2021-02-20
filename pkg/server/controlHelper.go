package server

import (
	"encoding/json"
	"log"
)

type jsonCommand struct {
	ClientID string
	Command  string
	Args     []string
}

func newCommand(command string) jsonCommand {
	return jsonCommand{
		Command: command,
	}
}

// func writeControlMessage(ws *websocket.Conn, clientID, cmd string, args ...string) {
// 	cmdStruct := newCommand(cmd)
// 	cmdStruct.ClientID = clientID
// 	if args != nil {
// 		for _, arg := range args {
// 			cmdStruct.Args = append(cmdStruct.Args, arg)
// 		}
// 	}
// 	rawMessage := structToByte(cmdStruct)
// 	ws.WriteMessage(websocket.TextMessage, rawMessage)
// }

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
func structToByte(cmd interface{}) []byte {
	rawMessage, err := json.Marshal(cmd)
	if err != nil {
		log.Panicf("Could not marshal: %v", err)
	}
	return rawMessage

}
