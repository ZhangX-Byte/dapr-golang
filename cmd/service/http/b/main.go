package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Zhang-Byte/dapr-golang/internal"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	"log"
	"net/http"
)
import daprd "github.com/dapr/go-sdk/service/http"

func main() {
	s := daprd.NewService(":9002")
	if err := s.AddServiceInvocationHandler("/hello", helloHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func helloHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}
	log.Printf("The go-service-b service method hello has been invoked,recieve message is %v", string(in.Data))

	msg := invokeService(ctx)
	httpResult := internal.HttpResult{Message: msg}

	out = &common.Content{
		Data:        httpResult.ToBytes(),
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}

func invokeService(ctx context.Context) (msg string) {
	client := client()
	content := &dapr.DataContent{
		ContentType: "text/plain",
		Data:        []byte("This is golang Service B."),
	}
	resp, err := client.InvokeServiceWithContent(ctx, "go-service-c", "hello", content)
	if err != nil {
		panic(err)
	}
	var result internal.HttpResult
	if err := json.Unmarshal(resp, &result); err != nil {
		log.Printf(err.Error())
	}
	msg = result.Message
	return
}

func client() dapr.Client {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	return client
}
