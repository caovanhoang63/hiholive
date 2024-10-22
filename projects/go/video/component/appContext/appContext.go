package appContext

import (
	"hiholive/shared/go/core"
	"hiholive/shared/go/srvctx"
)

type AppContext interface {
	core.AppContext
}

type appContext struct {
	logger srvctx.Logger
}

func NewAppContext(logger srvctx.Logger) *appContext {
	return &appContext{
		logger: logger,
	}
}

func (a *appContext) GetLogger() srvctx.Logger {
	return a.logger
}
