package web

import "servicelocator"

type WebHostBuilder struct {
	services      servicelocator.ServiceCollection
	configureFunc func(hosting *WebHost, services servicelocator.ServiceScope)
}

func (builder *WebHostBuilder) ConfigureServices(f func(servicelocator.ServiceCollection)) {
	services := servicelocator.NewServiceCollection()
	f(services)
}

func (builder *WebHostBuilder) Configure(f func(hosting *WebHost, services servicelocator.ServiceScope)) {
	builder.configureFunc = f
}

// func (builder *WebHostBuilder) Build() *WebHost {
// 	scope := builder.services.Build()
// 	hosting := &WebHost{
// 		rootScope: scope,
// 	}
// 	builder.configureFunc(hosting, scope)
// 	return hosting
// }
