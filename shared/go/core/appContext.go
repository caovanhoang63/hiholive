package core

import (
	"hiholive/shared/go/srvctx/components/loggerc"
)

type AppContext interface {
	GetLogger() loggerc.Logger
}
