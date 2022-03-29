package httpabstracts

import servicelocator "github.com/ah-its-andy/downton/serviceLocator"

type Startup interface {
	ConfigureServices(services servicelocator.ServiceCollection)
	CreateServiceScope(services servicelocator.ServiceCollection) servicelocator.ServiceScope
	Configure(app AppBuilder)
}
