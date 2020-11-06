package main

import (
	"context"
	dapr "github.com/dapr/go-sdk/client"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	// create the client
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	content := &dapr.DataContent{
		ContentType: "text/plain",
		Data:        []byte("This is client A."),
	}
	for {
		resp, err := client.InvokeServiceWithContent(ctx, "go-service-b", "hello", content)
		if err != nil {
			panic(err)
		}
		log.Printf("go-service-b method hello has invoked, response: %s", string(resp))

		time.Sleep(time.Second * 5)
	}

}
