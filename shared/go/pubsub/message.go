package pubsub

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Id        string    `json:"id"`
	Channel   string    `json:"channel"`
	Data      []byte    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()
	jsonData, _ := json.Marshal(data)

	return &Message{
		Id:        fmt.Sprintf("%d", now.UnixNano()),
		Data:      jsonData,
		CreatedAt: now,
	}
}

func (Message) String() string {

	return "Message"
}

func (m *Message) SetChannel(topic string) {
	m.Channel = topic
}

func (evt *Message) Marshal() ([]byte, error) {
	b, err := json.Marshal(evt)
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
