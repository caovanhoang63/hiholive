package common

type Config interface {
	GetGRPCPort() int
	GetGRPCAuthServerAddress() string
}
