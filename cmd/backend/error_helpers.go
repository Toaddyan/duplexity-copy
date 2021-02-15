package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func logInternalError(err error) {
	log.Panic(err)
}

func buildUnimplemntedError(funcName string) (err error) {
	return status.Errorf(codes.Unimplemented, fmt.Sprintf("method %s not implemented", funcName))
}

func buildInternalError(funcName string, recovery interface{}, stack string) (err error) {
	return status.Errorf(codes.Internal, fmt.Sprintf("backend recovering from panic in %s: %v\n%v", funcName, recovery, stack))
}

func buildNotFoundError(funcName string) (err error) {
	return status.Errorf(codes.NotFound, fmt.Sprintf("method %s requested an entity which could not be found", funcName))
}
