package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func Handler() (Response, error) {
	return Response{
		Message: "I'm not a APIGW interfacing lambda! Invoke me locally: sam local invoke Vanilla -e event.json",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
