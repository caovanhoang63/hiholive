package core

import "github.com/caovanhoang63/hiholive/shared/golang/srvctx"

type AppContext interface {
	GetLogger() srvctx.Logger
}
