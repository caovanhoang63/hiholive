package core

import (
	"hiholive/shared/go/srvctx"
)

type AppContext interface {
	GetLogger() srvctx.Logger
}
