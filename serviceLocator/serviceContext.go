package servicelocator

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ah-its-andy/downton/core"
)

type rootSvcCtx struct {
	services  map[string]*core.ServiceInfo
	instances map[string]any
}

func newRootSvcCtx(services map[string]*core.ServiceInfo) *rootSvcCtx {
	ctx := &rootSvcCtx{
		services:  make(map[string]*core.ServiceInfo),
		instances: make(map[string]any),
	}
	for _, s := range services {
		ctx.services[s.ServiceName] = s
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
		elemType := getElemType(s.ServiceType)
		instance := reflect.New(elemType)
		interfaceInstance := instance.Interface()
		ctx.instances[name] = interfaceInstance
	}
	return nil
}

func (ctx *rootSvcCtx) initReferences() error {
	for name, instance := range ctx.instances {
		serviceInfo := ctx.services[name]
		for _, ref := range serviceInfo.References {
			elemValue := getElemValue(reflect.ValueOf(instance))
			field := elemValue.FieldByName(ref.FieldName)
			if !field.IsValid() {
				return errors.New(fmt.Sprintf("%s has no field %s", name, ref.FieldName))
			}
			if referenceInstance, ok := ctx.instances[ref.ReferenceName]; ok {
				field.Set(reflect.ValueOf(referenceInstance))
				continue
			}

			if _, ok := ctx.services[ref.ReferenceName]; ok {
				return errors.New(fmt.Sprintf("%s depends on %s, which is not singleton", name, ref.ReferenceName))
			} else {
				return errors.New(fmt.Sprintf("%s depends on %s, which is not registered", name, ref.ReferenceName))
			}
		}
	}
	return nil
}

type scopeCtx struct {
	services           map[string]*core.ServiceInfo
	signletonInstances map[string]any
	scopeInstances     map[string]any
	parent             core.ServiceContext
}

func newScopeCtx(parent core.ServiceContext, services map[string]*core.ServiceInfo, instances map[string]any) *scopeCtx {
	ctx := &scopeCtx{
		services:           make(map[string]*core.ServiceInfo),
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
		for _, ref := range serviceInfo.References {
			elemValue := getElemValue(reflect.ValueOf(instance))
			field := elemValue.FieldByName(ref.FieldName)
			if !field.IsValid() {
				return errors.New(fmt.Sprintf("%s has no field %s", name, ref.FieldName))
			}
			if referenceInstance, ok := ctx.signletonInstances[ref.ReferenceName]; ok {
				field.Set(reflect.ValueOf(referenceInstance))
				continue
			}
			if referenceInstance, ok := ctx.scopeInstances[ref.ReferenceName]; ok {
				field.Set(reflect.ValueOf(referenceInstance))
				continue
			}

			if _, ok := ctx.services[ref.ReferenceName]; !ok {
				return errors.New(fmt.Sprintf("%s depends on %s, which is not registered", name, ref.ReferenceName))
			} else {
				// here service only can be transient
				return errors.New(fmt.Sprintf("%s depends on %s, which is transient", name, ref.ReferenceName))
			}
		}
	}
	return nil
}

func (ctx *scopeCtx) CreateTransientService(serviceInfo *core.ServiceInfo) (any, error) {
	transientInstance := reflect.New(serviceInfo.ServiceType).Interface()
	for _, transientRef := range serviceInfo.References {
		elemValue := getElemValue(reflect.ValueOf(transientInstance))
		field := elemValue.FieldByName(transientRef.FieldName)
		if !field.IsValid() {
			return nil, errors.New(fmt.Sprintf("%s has no field %s", serviceInfo.ServiceName, transientRef.FieldName))
		}
		if referenceInstance, ok := ctx.signletonInstances[transientRef.ReferenceName]; ok {
			field.Set(reflect.ValueOf(referenceInstance))
			continue
		}
		if referenceInstance, ok := ctx.scopeInstances[transientRef.ReferenceName]; ok {
			field.Set(reflect.ValueOf(referenceInstance))
			continue
		}
		if dependsServiceInfo, ok := ctx.services[transientRef.ReferenceName]; !ok {
			return nil, errors.New(fmt.Sprintf("%s depends on %s, which is not registered", serviceInfo.ServiceName, transientRef.ReferenceName))
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

func getElemType(t reflect.Type) reflect.Type {
	if t.Kind() != reflect.Ptr {
		return t
	}
	return getElemType(t.Elem())
}

func getElemValue(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		return v
	}
	return getElemValue(v.Elem())
}
