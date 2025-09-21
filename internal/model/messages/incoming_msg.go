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
	return s.tgClient.SendMessage(msg.Text, msg.UserId)
}
