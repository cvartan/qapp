package message

import "context"

// Base type for data transition
type Message struct {
	ctx     context.Context
	headers map[string]interface{}
	data    interface{}
}

// Create new message with base context
func New() *Message {
	return &Message{
		ctx:     context.Background(),
		headers: make(map[string]interface{}),
	}
}

// Create new message with user context
func NewWithContext(ctx context.Context) *Message {
	return &Message{
		ctx:     ctx,
		headers: make(map[string]interface{}),
	}
}

func (m *Message) SetData(data interface{}) {
	m.data = data
}

func (m *Message) GetData() interface{} {
	return m.data
}

func (m *Message) SetHeader(name string, value interface{}) {
	if name == "" {
		panic("name will not be empty")
	}
	m.headers[name] = value
}

func (m *Message) GetHeader(name string) (interface{}, bool) {
	if name == "" {
		panic("name will not be empty ")
	}
	value, ok := m.headers[name]
	return value, ok
}

func (m *Message) GetContext() context.Context {
	return m.ctx
}

func (m *Message) SetContext(ctx context.Context) {
	m.ctx = ctx
}

// Interface for realization request-response operations
type RequestProvider interface {
	Request(request Message, response *Message) error
}

// Interface for realization sending message operation
type MessageSender interface {
	Send(message Message) error
}

// Function for handle messages
type MessageHandler func(message Message)

// Interface for realization message reader
type MessageReader interface {
	Read(handler MessageHandler) error
}
