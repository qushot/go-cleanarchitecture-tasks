package viewmodels

type PrintViewModel struct {
	messages     []string
	messageColor Color
}

func NewPrintViewModel(messages []string, messageColor Color) *PrintViewModel {
	return &PrintViewModel{messages: messages, messageColor: messageColor}
}

func (p *PrintViewModel) GetMessages() []string {
	return p.messages
}

func (p *PrintViewModel) GetMessageColor() Color {
	return p.messageColor
}
