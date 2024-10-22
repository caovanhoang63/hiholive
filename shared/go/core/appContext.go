package core

import "github.com/caovanhoang63/hiholive/shared/srvctx"

type AppContext interface {
	GetLogger() srvctx.Logger
}
