package rabbitpubsub

import (
	"context"
	"flag"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/pubsub"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type rabbitPubSub struct {
	channel *amqp.Channel
	locker  *sync.RWMutex
	id      string
	logger  srvctx.Logger
	dsn     string
}

func (pb *rabbitPubSub) ID() string {
	return pb.id
}

func (pb *rabbitPubSub) InitFlags() {
	flag.StringVar(&pb.dsn, "rabbitmq-dsn", "", "rabbitmq connection string")
}

func (pb *rabbitPubSub) Activate(serviceContext srvctx.ServiceContext) error {
	rabbitConn, err := amqp.Dial(pb.dsn)

	logger := serviceContext.Logger(pb.id)

	if err != nil {
		return err
	}

	rabbitChannel, err := rabbitConn.Channel()
	if err != nil {
		return err
	}
	pb.channel = rabbitChannel
	logger.Info("Success to connect RabbitMQ")
	return nil
}

func (pb *rabbitPubSub) Stop() error {
	err := pb.channel.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewRabbitPubSub(id string) *rabbitPubSub {
	return &rabbitPubSub{
		id:     id,
		locker: new(sync.RWMutex),
	}
}

func (pb *rabbitPubSub) Publish(ctx context.Context, exchange string, data *pubsub.Message) error {
	go func() {
		defer core.AppRecover()

		data.SetTopic(exchange)

		bdata, err := data.Marshal()
		if err != nil {
			panic(err)
		}

		err = pb.channel.ExchangeDeclare(
			"logs",   // name
			"fanout", // type
			true,     // durable
			false,    // auto-deleted
			false,    // internal
			false,    // no-wait
			nil,      // arguments
		)
		failOnError(err, "Failed to declare an exchange")

		err = pb.channel.PublishWithContext(
			ctx,
			exchange,
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        bdata,
			})
		if err != nil {
			panic(err)
		}
		log.Println("New event published:", data, "with data", data.Data)
	}()
	return nil
}

func (pb *rabbitPubSub) Subscribe(ctx context.Context, exchange string) (<-chan *pubsub.Message, func()) {
	err := pb.channel.ExchangeDeclare(
		exchange, // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	q, err := pb.channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = pb.channel.QueueBind(
		q.Name,   // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := pb.channel.Consume(
		"",    // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to register a consumer")
	ch := make(chan *pubsub.Message)

	// TODO: handle this error
	go func() {
		for d := range msgs {
			var mess pubsub.Message
			mess.Unmarshal(d.Body)
			ch <- &mess
		}
	}()

	return ch, func() {
		return
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
