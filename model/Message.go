package model

type Message struct {
	From    string
	To      string
	Content string
}

func NewMessage(from, to, content string) *Message {
	return &Message{
		From:    from,
		To:      to,
		Content: content,
	}
}

func (m *Message) GetFrom() string {
	return m.From
}

func (m *Message) SetFrom(from string) {
	m.From = from
}

func (m *Message) GetTo() string {
	return m.To
}

func (m *Message) SetTo(to string) {
	m.To = to
}

func (m *Message) GetContent() string {
	return m.Content
}

func (m *Message) SetContent(content string) {
	m.Content = content
}
