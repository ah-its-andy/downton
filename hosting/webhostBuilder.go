package hosting

import (
	"github.com/ah-its-andy/downton/core"
)

type WebHostBuilder interface {
	UseEnvironment(env string) WebHostBuilder
	ConfigureServices(func(core.ServiceCollection)) WebHostBuilder
}
