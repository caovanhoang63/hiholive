package redisc

import (
	"flag"
	"fmt"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type RedisOption struct {
	dsn string
}

type redisEngine struct {
	id     string
	prefix string
	logger srvctx.Logger
	client *redis.Client
	*RedisOption
}

func NewRedis(id string) *redisEngine {
	return &redisEngine{
		RedisOption: new(RedisOption),
		id:          id,
	}
}
func (r *redisEngine) ID() string {
	return r.id
}

func (r *redisEngine) InitFlags() {
	prefix := r.prefix
	if r.prefix != "" {
		prefix += "-"
	}
	flag.StringVar(
		&r.dsn,
		fmt.Sprintf("%sredis-dsn", prefix),
		"",
		"Redis dsn",
	)

}

func (r *redisEngine) GetClient() *redis.Client {
	return r.client
}

func (r *redisEngine) Activate(ctx srvctx.ServiceContext) error {
	r.logger = srvctx.GlobalLogger().GetLogger(r.id)

	r.logger.Info("Connecting to redis...")

	opts, err := redis.ParseURL(r.dsn)
	if err != nil {
		panic(err)
	}
	r.client = redis.NewClient(opts)
	if err = r.client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	r.logger.Info("Connected to redis")

	return nil
}

func (r *redisEngine) Stop() error {
	return nil
}
