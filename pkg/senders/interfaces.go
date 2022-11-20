package senders

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// source: https://medium.com/@matryer/meet-moq-easily-mock-interfaces-in-go-476444187d10

//go:generate moq -out ../../mocks/email_sender.go -pkg mocks . EmailSender
type EmailSender interface {
	Send(email string) error
}

//go:generate moq -out ../../mocks/file_sender.go -pkg mocks . FileSender
type FileSender interface {
	Discover() ([]string, error)
}

//go:generate moq -out ../../mocks/dynamo_db.go -pkg mocks . DynamoDB
type DynamoDB interface {
	dynamodbiface.DynamoDBAPI
}
