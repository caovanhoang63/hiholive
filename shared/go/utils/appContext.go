package utils

import "hiholive/shared/go/logger"

type AppContext interface {
	GetLogger() logger.Logger
}
