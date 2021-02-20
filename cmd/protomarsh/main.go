package main

import (
	"log"

	wspb "github.com/duplexityio/duplexity/cmd/protomarsh/wspb/websocket/v1"
	"github.com/golang/protobuf/proto"
)

func main() {
	foo := &wspb.ClientRequest{
		Type: &wspb.ClientRequest_RegisterConnectionRequest{
			RegisterConnectionRequest: &wspb.RegisterConnectionRequest{
				Client: &wspb.Client{
					ClientId: "here is a clientid",
				},
			},
		},
	}

	rawBytes, err := proto.Marshal(foo)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(string(rawBytes))

	bar := &wspb.ClientRequest{}
	err = proto.Unmarshal(rawBytes, bar)
	log.Printf("%+v\n", bar)

	wrapped := bar.GetType()
	log.Println("")

	/*
	   {
	   clientId: "clientIDGoesHere",
	   eventType: "registerConnection",
	   fields: ["", "", ""],
	   }
	*/
}
