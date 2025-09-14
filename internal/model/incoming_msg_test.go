package model

import (
	mock_msg "route256/internal/mock/message"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_IncommingMessage_Send_Hello_Message(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	modelMock := mock_msg.NewMockMessageSender(ctrl)
	model := New(modelMock)
	modelMock.EXPECT().SendMessage("Привет", 1).Return(nil)

	err := model.IncommingMessage(Message{TextCommand: "/start", UserId: 1})

	assert.NoError(t, err)
}

func Test_IncommingMessage_Send_Hello_Not_Know_Message(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	modelMock := mock_msg.NewMockMessageSender(ctrl)
	model := New(modelMock)
	modelMock.EXPECT().SendMessage("Пока не знаю такой команды", 1).Return(nil)

	err := model.IncommingMessage(Message{TextCommand: "/dwwaffwfgwg", UserId: 1})

	assert.NoError(t, err)
}
