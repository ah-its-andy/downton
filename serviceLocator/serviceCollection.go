package servicelocator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ServiceCollection interface {
	AddService(serviceInfo *ServiceInfo)
	AddLifetimeService(service any, lifetime int)
	AddSingleton(service any)
	AddScope(service any)
	AddTransient(service any)

	Build() ServiceScope
}

type ServiceInfo struct {
	ServiceType reflect.Type
	Lifetime    int

	references []*ServiceReference
}

type ServiceReference struct {
	ReferenceType reflect.Type
	FieldName     string
}

type svcCollection struct {
	services map[string]*ServiceInfo
}

func NewServiceCollection() ServiceCollection {
	return &svcCollection{
		services: make(map[string]*ServiceInfo),
	}
}

func (services *svcCollection) AddService(serviceInfo *ServiceInfo) {
	serviceName := getTypeFullName(serviceInfo.ServiceType)
	if serviceInfo.ServiceType.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("Service type must be a pointer type, but %s is %s", serviceName, serviceInfo.ServiceType.String()))
	}

	if _, ok := services.services[serviceName]; ok {
		panic(fmt.Sprintf("service %s already registered", serviceName))
	}
	services.services[serviceName] = serviceInfo
}

func (services *svcCollection) AddLifetimeService(service any, lifetime int) {
	serviceInfo := &ServiceInfo{
		ServiceType: reflect.TypeOf(service),
		Lifetime:    lifetime,
	}
	services.AddService(serviceInfo)
}

func (services *svcCollection) AddSingleton(service any) {
	services.AddLifetimeService(service, LifetimeSingleton)
}

func (services *svcCollection) AddScope(service any) {
	services.AddLifetimeService(service, LifetimeScope)
}

func (services *svcCollection) AddTransient(service any) {
	services.AddLifetimeService(service, LifetimeTransient)
}

func (services *svcCollection) buildServiceDependencies() {
	for _, serviceInfo := range services.services {
		resolveServiceDependencies(serviceInfo)
	}
}

func (services *svcCollection) checkCycleReferences(serviceName string, references []*ServiceReference, path string) error {
	pathList := make([]string, 0)
	if path != "" {
		pathList = append(pathList, path)
	}
	pathList = append(pathList, serviceName)
	curPath := strings.Join(pathList, " -> ")
	for _, ref := range references {
		dependsServiceName := getTypeFullName(ref.ReferenceType)
		if serviceName == dependsServiceName {
			return errors.New(fmt.Sprintf("%s has a cycle reference to %s", curPath, dependsServiceName))
		}
		if dependsServiceInfo, ok := services.services[dependsServiceName]; ok {
			err := services.checkCycleReferences(serviceName, dependsServiceInfo.references, curPath)
			if err != nil {
				return err
			}
		} else {
			return errors.New(fmt.Sprintf("%s depends on %s, which is not registered", curPath, dependsServiceName))
		}
	}
	return nil
}

func (services *svcCollection) Build() ServiceScope {
	services.buildServiceDependencies()
	for serviceName, serviceInfo := range services.services {
		err := services.checkCycleReferences(serviceName, serviceInfo.references, "")
		if err != nil {
			panic(err)
		}
	}
	return &rootSvcScope{
		rootSvcCtx: *newRootSvcCtx(services.services),
	}
}

func resolveServiceDependencies(serviceInfo *ServiceInfo) {
	elemType := getElemType(serviceInfo.ServiceType)
	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		if field.Type.Kind() != reflect.Ptr {
			continue
		}
		if field.Tag.Get("inject") == "ignore" {
			continue
		}
		serviceInfo.references = append(serviceInfo.references, &ServiceReference{
			ReferenceType: field.Type,
			FieldName:     field.Name,
		})
	}
}
