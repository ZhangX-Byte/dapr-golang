package main

import (
	"context"
	"errors"
	"github.com/dapr/go-sdk/service/common"
	"log"
	"net/http"
)
import daprd "github.com/dapr/go-sdk/service/http"

func main() {
	s := daprd.NewService(":8080")
	if err := s.AddServiceInvocationHandler("/echo", echoHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func echoHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}
	log.Printf(
		"echo - ContentType:%s, Verb:%s, QueryString:%s, %s",
		in.ContentType, in.Verb, in.QueryString, in.Data,
	)
	out = &common.Content{
		Data:        in.Data,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}
