package messages

import (
	"reflect"
)

// controlMessageTypeRegister is a map of control message struct names mapping to their reflect.Type
var controlMessageTypeRegister map[string]reflect.Type

// registerControlMessageType registers a control message into the ControlMessageTypeRegister
func registerControlMessageType(controlMessage interface{}) {
	controlMessageTypeRegister[reflect.TypeOf(controlMessage).Name()] = reflect.TypeOf(controlMessage)
}

func init() {
	controlMessageTypeRegister = make(map[string]reflect.Type)
	registerControlMessageType(DiscoveryRequest{})
	registerControlMessageType(DiscoveryResponse{})
	registerControlMessageType(PipesRegistered{})
}
