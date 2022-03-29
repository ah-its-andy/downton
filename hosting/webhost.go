package hosting

import servicelocator "github.com/ah-its-andy/downton/serviceLocator"

type WebHost interface {
	GetServiceScope() servicelocator.ServiceScope

	Start()
	Stop()
}
