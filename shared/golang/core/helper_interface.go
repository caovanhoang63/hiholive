package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type GINComponent interface {
	GetPort() int
	GetRouter() *gin.Engine
}

type GormComponent interface {
	GetDB() *gorm.DB
}

type RedisComponent interface {
	GetClient() *redis.Client
}

type JWTProvider interface {
	IssueToken(ctx context.Context, id, sub string) (token string, expSecs int, err error)
	ParseToken(ctx context.Context, tokenString string) (claims *jwt.RegisteredClaims, err error)
}

type Config interface {
	GetGRPCPort() int
	GetGRPCServerAddress() string

	GetGRPCUserAddress() string
	GetGRPCAuthAddress() string
	GetGRPCHlsAddress() string
	GetGRPCRtmpAddress() string
	GetGRPCAnalyticAddress() string
	GetGRPCCommunicationAddress() string
	GetGRPCVideoAddress() string
}
