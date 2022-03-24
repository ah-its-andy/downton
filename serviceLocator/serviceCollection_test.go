package servicelocator

import (
	"reflect"
	"testing"
)

type ServiceCollectionTestSingletonService struct {
}

type ServiceCollectionTestScopeService struct {
}

type ServiceCollectionTestTransientService struct {
}

type ServiceCollectionTestReferenceService struct {
	NotService   string
	ShouldIgnore *ServiceCollectionTestSingletonService `inject:"ignore"`
	NotPtr       ServiceCollectionTestSingletonService
	ShouldInject *ServiceCollectionTestSingletonService
}

func Test_ServiceCollection_Constructor(t *testing.T) {
	services := NewServiceCollection()
	if services == nil {
		t.Error("services is nil")
	}
}

func TestServiceCollectionAddService(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("panic")
		}
	}()

	serviceInfo := &ServiceInfo{
		ServiceType: reflect.TypeOf(&ServiceCollectionTestSingletonService{}),
		Lifetime:    LifetimeSingleton,
	}
	services := NewServiceCollection()
	services.AddService(serviceInfo)
}

func TestServiceCollectionAddServiceNotPointer(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect panic, but passed")
		}
	}()

	serviceInfo := &ServiceInfo{
		ServiceType: reflect.TypeOf(ServiceCollectionTestSingletonService{}),
		Lifetime:    LifetimeSingleton,
	}
	services := NewServiceCollection()
	services.AddService(serviceInfo)
}

func TestServiceCollectionAddServiceDeplicate(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect panic, but passed")
		}
	}()
	serviceInfo := &ServiceInfo{
		ServiceType: reflect.TypeOf(ServiceCollectionTestSingletonService{}),
		Lifetime:    LifetimeSingleton,
	}
	services := NewServiceCollection()
	services.AddService(serviceInfo)
	services.AddService(serviceInfo)
}

func Test_ServiceCollection_AddLifetimeService(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	services := NewServiceCollection()
	services.AddLifetimeService(&ServiceCollectionTestSingletonService{}, LifetimeSingleton)
}

func Test_ServiceCollection_AddLifetimeService_NotPointer(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect panic, but passed")
		}
	}()
	services := NewServiceCollection()
	services.AddLifetimeService(ServiceCollectionTestSingletonService{}, LifetimeSingleton)
}

func Test_ServiceCollection_AddLifetimeService_Duplicated(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect panic, but passed")
		}
	}()
	services := NewServiceCollection()
	services.AddLifetimeService(ServiceCollectionTestSingletonService{}, LifetimeSingleton)
	services.AddLifetimeService(ServiceCollectionTestSingletonService{}, LifetimeSingleton)
}

func Test_ServiceCollection_AddSingleton(t *testing.T) {
	services := NewServiceCollection()
	services.AddSingleton(&ServiceCollectionTestSingletonService{})
}

func Test_ServiceCollection_AddScope(t *testing.T) {
	services := NewServiceCollection()
	services.AddScope(&ServiceCollectionTestScopeService{})
}

func Test_ServiceCollection_AddTransient(t *testing.T) {
	services := NewServiceCollection()
	services.AddTransient(&ServiceCollectionTestTransientService{})
}

func Test_ServiceCollection_ResolveDependencies(t *testing.T) {
	services := NewServiceCollection()
	services.AddSingleton(&ServiceCollectionTestSingletonService{})

	serviceInfo := &ServiceInfo{
		ServiceType: reflect.TypeOf(&ServiceCollectionTestReferenceService{}),
		Lifetime:    LifetimeSingleton,
	}
	services.AddService(serviceInfo)
	services.Build()

	shouldInject := false
	for _, depends := range serviceInfo.references {
		if depends.FieldName == "ShouldIgnore" {
			t.Error("ShouldIgnore should be ignored")
		} else if depends.FieldName == "NotPtr" {
			t.Error("NotPtr should be ignored")
		} else if depends.FieldName == "NotService" {
			t.Error("NotService should be ignored")
		} else if depends.FieldName == "ShouldInject" {
			shouldInject = true
		}
	}
	if !shouldInject {
		t.Error("shouldInject should be resolved")
	}
}
