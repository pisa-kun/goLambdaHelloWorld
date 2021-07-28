package main

import (
	"TestLambda/greeting"

	"github.com/aws/aws-lambda-go/lambda"
)

func excuteFunction() {
	greeting.SayHello()
}

func main() {
	lambda.Start(excuteFunction)
}
