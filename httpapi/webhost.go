package httpapi

import (
	"github.com/ah-its-andy/downton/core"
	"github.com/ah-its-andy/downton/hosting"
)

type WebHost interface {
	hosting.Host

	ServiceScope() core.ServiceScope
	Start()
	Stop()
}
