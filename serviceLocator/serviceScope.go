package servicelocator

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ah-its-andy/downton/core"
)

type rootSvcScope struct {
	rootSvcCtx
}

func (scope *rootSvcScope) CreateScope() core.ServiceScope {
	ctx := newScopeCtx(&scope.rootSvcCtx, scope.services, scope.instances)
	return &scopedSvcScope{
		scopeCtx: *ctx,
	}
}

func (scope *rootSvcScope) GetService(service any) any {
	svc, err := scope.GetRequiredService(service)
	if err != nil {
		return nil
	}
	return svc
}
func (scope *rootSvcScope) GetRequiredService(service any) (any, error) {
	serviceType := reflect.TypeOf(service)
	serviceName := getTypeFullName(serviceType)
	if service, ok := scope.instances[serviceName]; ok {
		return service, nil
	}

	if serviceInfo, ok := scope.services[serviceName]; ok {
		return nil, errors.New(fmt.Sprintf("only singleton service provided by root scope, but %s is %s", serviceName, GetLifetimeName(serviceInfo.Lifetime)))
	} else {
		return nil, errors.New(fmt.Sprintf("service %s not found", serviceName))
	}
}

func (scope *rootSvcScope) GetServices(service any) []any {
	return GetServices(scope, scope.services, service)
}

func GetServices(scope core.ServiceScope, serviceInfos map[string]*core.ServiceInfo, service any) []any {
	t := reflect.TypeOf(service)
	if t.Kind() != reflect.Interface {
		panic(fmt.Sprintf("ServiceScope.GetServices only accept interface, but %s is %s", t.Name(), t.Kind()))
	}
	results := make([]any, 0)
	for _, serviceInfo := range serviceInfos {
		if serviceInfo.ServiceType.Implements(t) {
			svc := scope.GetService(serviceInfo.ServiceType)
			if svc != nil {
				results = append(results, svc)
			}
		}
	}
	return results
}

type scopedSvcScope struct {
	scopeCtx
}

func (scope *scopedSvcScope) CreateScope() core.ServiceScope {
	ctx := newScopeCtx(&scope.scopeCtx, scope.services, scope.signletonInstances)
	return &scopedSvcScope{
		scopeCtx: *ctx,
	}
}

func (scope *scopedSvcScope) GetRequiredService(service any) (any, error) {
	serviceType := reflect.TypeOf(service)
	serviceName := getTypeFullName(serviceType)
	if service, ok := scope.signletonInstances[serviceName]; ok {
		return service, nil
	}

	if service, ok := scope.scopeInstances[serviceName]; ok {
		return service, nil
	}

	if serviceInfo, ok := scope.services[serviceName]; ok {
		return scope.CreateTransientService(serviceInfo)
	} else {
		return nil, errors.New(fmt.Sprintf("service %s not found", serviceName))
	}
}
func (scope *scopedSvcScope) GetService(service any) any {
	svc, err := scope.GetRequiredService(service)
	if err != nil {
		return nil
	}
	return svc
}

func (scope *scopedSvcScope) GetServices(service any) []any {
	return GetServices(scope, scope.services, service)
}
