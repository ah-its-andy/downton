package httpapi

import (
	"github.com/ah-its-andy/downton/hosting"
	servicelocator "github.com/ah-its-andy/downton/serviceLocator"
)

type WebHost interface {
	hosting.Host

	ServiceScope() servicelocator.ServiceScope
	Start()
	Stop()
}
