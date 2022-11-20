package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mauleyzaola/logtests/pkg/senders"
)

func main() {}

func Send(sender senders.EmailSender, email string) error {
	return sender.Send(email)
}

func CreateDBBackup(db senders.DynamoDB, backupName, tableName string) error {
	input := &dynamodb.CreateBackupInput{
		BackupName: aws.String(backupName),
		TableName:  aws.String(tableName),
	}
	_, err := db.CreateBackup(input)
	return err
}
