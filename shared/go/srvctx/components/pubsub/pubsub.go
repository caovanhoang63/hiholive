package pubsub

import "context"

type Pubsub interface {
	Publish(ctx context.Context, topic string, data *Message) error
	Subscribe(ctx context.Context, topic string) (ch <-chan *Message, close func())
}
