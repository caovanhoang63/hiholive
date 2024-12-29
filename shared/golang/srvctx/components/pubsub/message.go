package pubsub

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Id        string    `json:"id"`
	Channel   string    `json:"channel"`
	Data      any       `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()
	id, _ := uuid.NewUUID()
	return &Message{
		Id:        id.String(),
		Data:      data,
		CreatedAt: now,
	}
}

func (m *Message) SetTopic(topic string) {
	m.Channel = topic
}

func (m *Message) Marshal() ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (m *Message) Unmarshal(b []byte) error {
	err := json.Unmarshal(b, m)
	if err != nil {
		return err
	}
	return nil
}
