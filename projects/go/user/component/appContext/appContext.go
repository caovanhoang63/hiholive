package appContext

import (
	"hiholive/shared/go/logger"
	"hiholive/shared/go/utils"
)

type AppContext interface {
	utils.AppContext
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
