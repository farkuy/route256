package tg

import (
	"log"
	"log/slog"
	"route256/internal/model/messages"
	"route256/internal/model/spending"
	"strconv"
	"strings"

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

func (c *Client) ListenUpdates(msgModel *messages.Model, spendingModel *spending.SpendingsUsersStorage) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			reqText := update.Message.Text
			resTextMessage := ""
			var err error

			//TODO переписать ублюдство с рабиением этого на отдельные сегменты
			switch {
			case strings.HasPrefix(reqText, "/start"):
				resTextMessage = "Привет"
			case strings.HasPrefix(reqText, "/addSum"):
				spendingData, err := spending.ParseSendSpendingComand(reqText)
				if err != nil {
					resTextMessage = err.Error()
					break
				}
				err = spendingModel.Store.SendSpending(update.Message.From.ID, spendingData.Sum, spendingData.SpendingType, spendingData.Date)
				if err != nil {
					resTextMessage = err.Error()
					break
				}
				resTextMessage = "Трата добавлена"
			case strings.HasPrefix(reqText, "/getSpending"):
				period, err := spending.ParseGetUserSpendingHistory(reqText)
				if err != nil {
					resTextMessage = err.Error()
					break
				}
				data, err := spendingModel.Store.GetUserSpendingHistory(update.Message.From.ID, spending.TimePeriod(period))
				if err != nil {
					resTextMessage = err.Error()
					break
				}
				resTextMessage = "За " + period + " вы потратили на "
				for key, val := range data {
					sum := strconv.Itoa(val)
					resTextMessage += string(key) + " : " + sum + "; "
				}
			default:
				resTextMessage = "Я пока еще не знаю такую команду "
			}

			err = msgModel.IncommingMessage(messages.Message{
				Text:   resTextMessage,
				UserId: update.Message.From.ID,
			})
			if err != nil {
				errors.Wrap(err, "error processing message")
			}
		}
	}
}
