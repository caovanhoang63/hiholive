package core

import (
	"flag"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
)

type config struct {
	grpcPort          int    // for server port listening
	grpcServerAddress string // for client make grpc client connection

	grpcUserAddress          string
	grpcAuthAddress          string
	grpcHlsAddress           string
	grpcRtmpAddress          string
	grpcAnalyticAddress      string
	grpcCommunicationAddress string
	grpcVideoAddress         string
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ID() string {
	return KeyCompConf
}

func (c *config) InitFlags() {
	flag.IntVar(
		&c.grpcPort,
		"grpc-port",
		3100,
		"gRPC Port. Default: 3100",
	)

	flag.StringVar(
		&c.grpcServerAddress,
		"grpc-server-address",
		"localhost:3101",
		"gRPC server address. Default: localhost:3101",
	)

	flag.StringVar(
		&c.grpcUserAddress,
		"grpc-user-address",
		"localhost:3201",
		"gRPC user address. Default: localhost:3201",
	)

	flag.StringVar(
		&c.grpcAuthAddress,
		"grpc-auth-address",
		"localhost:3201",
		"gRPC user address. Default: localhost:3201",
	)

	flag.StringVar(
		&c.grpcHlsAddress,
		"grpc-hls-address",
		"localhost:3201",
		"gRPC user address. Default: localhost:3201",
	)

	flag.StringVar(
		&c.grpcRtmpAddress,
		"grpc-rtmp-address",
		"localhost:3201",
		"gRPC user address. Default: localhost:3201",
	)

	flag.StringVar(
		&c.grpcAnalyticAddress,
		"grpc-analytic-address",
		"localhost:3201",
		"gRPC user address. Default: localhost:3201",
	)

	flag.StringVar(
		&c.grpcVideoAddress,
		"grpc-video-address",
		"localhost:3201",
		"gRPC user address. Default: localhost:3201",
	)
	flag.StringVar(
		&c.grpcCommunicationAddress,
		"grpc-communication-address",
		"localhost:3201",
		"gRPC user address. Default: localhost:3201",
	)

}

func (c *config) Activate(_ srvctx.ServiceContext) error {
	return nil
}

func (c *config) Stop() error {
	return nil
}

func (c *config) GetGRPCPort() int {
	return c.grpcPort
}

func (c *config) GetGRPCServerAddress() string {
	return c.grpcServerAddress
}

func (c *config) GetGRPCUserAddress() string {
	return c.grpcUserAddress
}

func (c *config) GetGRPCAuthAddress() string {
	return c.grpcAuthAddress
}

func (c *config) GetGRPCHlsAddress() string {
	return c.grpcHlsAddress
}

func (c *config) GetGRPCRtmpAddress() string {
	return c.grpcRtmpAddress
}

func (c *config) GetGRPCAnalyticAddress() string {
	return c.grpcAnalyticAddress
}

func (c *config) GetGRPCCommunicationAddress() string {
	return c.grpcCommunicationAddress
}

func (c *config) GetGRPCVideoAddress() string {
	return c.grpcVideoAddress
}
