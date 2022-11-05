package main

import (
	"github.com/mauleyzaola/logtests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

const email = "test@example.com"

func TestSend_Moq(t *testing.T) {
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
