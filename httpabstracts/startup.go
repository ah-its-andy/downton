package httpabstracts

import (
	"github.com/ah-its-andy/downton/core"
)

type Startup interface {
	ConfigureServices(services core.ServiceCollection)
	CreateServiceScope(services core.ServiceCollection) core.ServiceScope
	Configure(app AppBuilder)
}
