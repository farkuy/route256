package messages

type MessageSender interface {
	SendMessage(message string, userId int64) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text   string
	UserId int64
}

func (s *Model) IncommingMessage(msg Message) error {
	if msg.Text == "/start" {
		return s.tgClient.SendMessage("Привет", msg.UserId)
	}

	return s.tgClient.SendMessage("Пока не знаю такой команды", msg.UserId)
}
