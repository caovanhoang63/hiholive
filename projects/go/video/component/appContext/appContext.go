package appContext

import (
	"hiholive/shared/go/core"
	"hiholive/shared/go/srvctx/components/loggerc"
)

type AppContext interface {
	core.AppContext
}

type appContext struct {
	logger loggerc.Logger
}

func NewAppContext(logger loggerc.Logger) *appContext {
	return &appContext{
		logger: logger,
	}
}

func (a *appContext) GetLogger() loggerc.Logger {
	return a.logger
}
