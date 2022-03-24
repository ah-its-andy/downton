package servicelocator

import (
	"errors"
	"fmt"
	"reflect"
)

type ServiceScope interface {
	GetRequiredService(service any) (any, error)
	GetService(service any) any

	CreateScope() ServiceScope
}

type rootSvcScope struct {
	rootSvcCtx
}

func (scope *rootSvcScope) CreateScope() ServiceScope {
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

type scopedSvcScope struct {
	scopeCtx
}

func (scope *scopedSvcScope) CreateScope() ServiceScope {
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
