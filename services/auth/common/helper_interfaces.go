package common

type Config interface {
	GetGRPCPort() int
	GetGRPCServerAddress() string
	GetGRPCUserAddress() string
}
