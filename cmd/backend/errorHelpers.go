package main

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func buildUnimplemntedError(funcName string) (err error) {
	err = status.Errorf(codes.Unimplemented, fmt.Sprintf("method %s not implemented", funcName))
	return
}
