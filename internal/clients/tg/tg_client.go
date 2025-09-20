package tg

import (
	"log"
	"log/slog"
	"route256/internal/model/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type Token interface {
	GetToken() string
}

type Client struct {
	client *tgbotapi.BotAPI
}

func Start(token Token) (*Client, error) {
	tgBot, err := tgbotapi.NewBotAPI(token.GetToken())
	if err != nil {
		return nil, errors.Wrap(err, "NewBotAPI")
	}

	slog.Info("Бот запустился")
	return &Client{client: tgBot}, nil
}

func (c *Client) SendMessage(message string, userId int64) error {
	_, err := c.client.Send(tgbotapi.NewMessage(userId, message))
	if err != nil {
		return errors.Wrap(err, "client send")
	}

	return nil
}

func (c *Client) ListenUpdates(msgModel *messages.Model) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			err := msgModel.IncommingMessage(messages.Message{
				Text:   update.Message.Text,
				UserId: update.Message.From.ID,
			})
			if err != nil {
				errors.Wrap(err, "error processing message")
			}
		}
	}
}
