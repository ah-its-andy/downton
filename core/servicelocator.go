package core

import "reflect"

type ServiceCollection interface {
	AddService(serviceInfo *ServiceInfo)
	AddLifetimeService(service any, lifetime int)
	AddSingleton(service any)
	AddScope(service any)
	AddTransient(service any)

	Build() ServiceScope
}

type ServiceInfo struct {
	ServiceName string
	ServiceType reflect.Type
	Lifetime    int

	References []*ServiceReference
}

type ServiceReference struct {
	ReferenceName string
	ReferenceType reflect.Type
	FieldName     string
}

type ServiceContext interface {
}

type ServiceScope interface {
	GetRequiredService(service any) (any, error)
	GetService(service any) any

	GetServices(service any) []any

	CreateScope() ServiceScope
}
