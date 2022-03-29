package hosting

import servicelocator "github.com/ah-its-andy/downton/serviceLocator"

type WebHostBuilder interface {
	UseEnvironment(env string) WebHostBuilder
	ConfigureServices(func(servicelocator.ServiceCollection)) WebHostBuilder
}
