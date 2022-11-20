package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mauleyzaola/logtests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

const email = "test@example.com"

func TestSend(t *testing.T) {
	var sentTo string
	mockedEmailSender := &mocks.EmailSenderMock{
		SendFunc: func(email string) error {
			sentTo = email
			return nil
		},
	}
	err := Send(mockedEmailSender, email)
	assert.NoError(t, err)
	assert.EqualValues(t, email, sentTo)
	expectedCount := 1
	assert.EqualValues(t, expectedCount, len(mockedEmailSender.SendCalls()))
}

func TestCreateDBBackup(t *testing.T) {
	var expectedBackupName, expectedTableName string
	mockedDynamoDB := &mocks.DynamoDBMock{
		CreateBackupFunc: func(input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
			expectedBackupName = *input.BackupName
			expectedTableName = *input.TableName
			return nil, fmt.Errorf("test error")
		},
	}
	err := CreateDBBackup(mockedDynamoDB, "backupName", "tableName")
	assert.Error(t, err)
	assert.EqualValues(t, expectedBackupName, "backupName")
	assert.EqualValues(t, expectedTableName, "tableName")
	expectedCount := 1
	assert.EqualValues(t, expectedCount, len(mockedDynamoDB.CreateBackupCalls()))
}
