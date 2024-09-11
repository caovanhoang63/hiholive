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

func NewAppContext(logger logger.Logger) *appContext {
	return &appContext{
		logger: logger,
	}
}

func (a *appContext) GetLogger() logger.Logger {
	return a.logger
}
