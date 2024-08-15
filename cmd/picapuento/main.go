package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Action string `json:"action"`
}

func Picapuento(ctx context.Context, event *Request) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	message := fmt.Sprintf("Action pretended is %s!", event.Action)
	return &message, nil
}

func main() {
	lambda.Start(Picapuento)
}
