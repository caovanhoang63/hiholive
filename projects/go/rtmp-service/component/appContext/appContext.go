package appContext

import (
	"hiholive/shared/go/common"
	"hiholive/shared/go/logger"
)

type AppContext interface {
	common.AppContext
}

type appContext struct {
	logger logger.Logger
}

func NewAppContextRTPM(logger logger.Logger) *appContext {
	return &appContext{
		logger: logger,
	}
}

func (a *appContext) GetLogger() logger.Logger {
	return a.logger
}

func (a *appContext) SetLogger(logger logger.Logger) {
	a.logger = logger
}
