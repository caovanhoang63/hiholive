package core

import "github.com/caovanhoang63/hiholive/shared/go/srvctx"

type AppContext interface {
	GetLogger() srvctx.Logger
}
