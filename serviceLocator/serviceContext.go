package servicelocator

import (
	"errors"
	"fmt"
	"reflect"
)

type ServiceContext interface {
}

type rootSvcCtx struct {
	services  map[string]*ServiceInfo
	instances map[string]any
}

func newRootSvcCtx(services map[string]*ServiceInfo) *rootSvcCtx {
	ctx := &rootSvcCtx{
		services:  make(map[string]*ServiceInfo),
		instances: make(map[string]any),
	}
	for _, s := range services {
		ctx.services[getTypeFullName(s.ServiceType)] = s
	}

	err := ctx.initServices()
	if err != nil {
		panic(err)
	}
	err = ctx.initReferences()
	if err != nil {
		panic(err)
	}
	return ctx
}

func (ctx *rootSvcCtx) initServices() error {
	for name, s := range ctx.services {
		if s.Lifetime != LifetimeSingleton {
			continue
		}
		instance := reflect.New(s.ServiceType).Interface()
		ctx.instances[name] = instance
	}
	return nil
}

func (ctx *rootSvcCtx) initReferences() error {
	for name, instance := range ctx.instances {
		serviceInfo := ctx.services[name]
		for _, ref := range serviceInfo.references {
			refFullName := getTypeFullName(ref.ReferenceType)
			field := reflect.ValueOf(instance).Elem().FieldByName(ref.FieldName)
			if !field.IsValid() {
				return errors.New(fmt.Sprintf("%s has no field %s", name, ref.FieldName))
			}
			if referenceInstance, ok := ctx.instances[refFullName]; ok {
				field.Set(reflect.ValueOf(referenceInstance))
				continue
			}

			if _, ok := ctx.services[refFullName]; ok {
				return errors.New(fmt.Sprintf("%s depends on %s, which is not singleton", name, refFullName))
			} else {
				return errors.New(fmt.Sprintf("%s depends on %s, which is not registered", name, refFullName))
			}
		}
	}
	return nil
}

type scopeCtx struct {
	services           map[string]*ServiceInfo
	signletonInstances map[string]any
	scopeInstances     map[string]any
	parent             ServiceContext
}

func newScopeCtx(parent ServiceContext, services map[string]*ServiceInfo, instances map[string]any) *scopeCtx {
	ctx := &scopeCtx{
		services:           make(map[string]*ServiceInfo),
		signletonInstances: make(map[string]any),
		scopeInstances:     make(map[string]any),
		parent:             parent,
	}
	for name, s := range services {
		ctx.services[name] = s
	}
	for name, instance := range instances {
		ctx.signletonInstances[name] = instance
	}
	return ctx
}

func (ctx *scopeCtx) initServices() error {
	for name, s := range ctx.services {
		if s.Lifetime != LifetimeScope {
			continue
		}
		instance := reflect.New(s.ServiceType).Interface()
		ctx.scopeInstances[name] = instance
	}
	return nil
}

func (ctx *scopeCtx) initReferences() error {
	for name, instance := range ctx.scopeInstances {
		serviceInfo := ctx.services[name]
		for _, ref := range serviceInfo.references {
			refFullName := getTypeFullName(ref.ReferenceType)
			field := reflect.ValueOf(instance).Elem().FieldByName(ref.FieldName)
			if !field.IsValid() {
				return errors.New(fmt.Sprintf("%s has no field %s", name, ref.FieldName))
			}
			if referenceInstance, ok := ctx.signletonInstances[refFullName]; ok {
				field.Set(reflect.ValueOf(referenceInstance))
				continue
			}
			if referenceInstance, ok := ctx.scopeInstances[refFullName]; ok {
				field.Set(reflect.ValueOf(referenceInstance))
				continue
			}

			if _, ok := ctx.services[refFullName]; !ok {
				return errors.New(fmt.Sprintf("%s depends on %s, which is not registered", name, refFullName))
			} else {
				// here service only can be transient
				return errors.New(fmt.Sprintf("%s depends on %s, which is transient", name, refFullName))
			}
		}
	}
	return nil
}

func (ctx *scopeCtx) CreateTransientService(serviceInfo *ServiceInfo) (any, error) {
	serviceName := getTypeFullName(serviceInfo.ServiceType)
	transientInstance := reflect.New(serviceInfo.ServiceType).Interface()
	for _, transientRef := range serviceInfo.references {
		refFullName := getTypeFullName(transientRef.ReferenceType)
		field := reflect.ValueOf(transientInstance).Elem().FieldByName(transientRef.FieldName)
		if !field.IsValid() {
			return nil, errors.New(fmt.Sprintf("%s has no field %s", serviceName, transientRef.FieldName))
		}
		if referenceInstance, ok := ctx.signletonInstances[refFullName]; ok {
			field.Set(reflect.ValueOf(referenceInstance))
			continue
		}
		if referenceInstance, ok := ctx.scopeInstances[refFullName]; ok {
			field.Set(reflect.ValueOf(referenceInstance))
			continue
		}
		if dependsServiceInfo, ok := ctx.services[refFullName]; !ok {
			return nil, errors.New(fmt.Sprintf("%s depends on %s, which is not registered", serviceName, refFullName))
		} else {
			dependsTransientInstance, err := ctx.CreateTransientService(dependsServiceInfo)
			if err != nil {
				return nil, err
			}
			field.Set(reflect.ValueOf(dependsTransientInstance))
		}
	}
	return transientInstance, nil
}

func getTypeFullName(t reflect.Type) string {
	v := t
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.PkgPath() + "." + v.Name()
}
