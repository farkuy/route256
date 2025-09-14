package model

type MessageSender interface {
	SendMessage(message string, userId int) error
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
	TextCommand string
	UserId      int
}

func (s *Model) IncommingMessage(msg Message) error {
	if msg.TextCommand == "/start" {
		s.tgClient.SendMessage("Привет", msg.UserId)
		return nil
	}

	s.tgClient.SendMessage("Пока не знаю такой команды", msg.UserId)
	return nil
}
