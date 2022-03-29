package config

import (
	"github.com/ah-its-andy/downton/collections"
	"github.com/ah-its-andy/downton/core"
)

type ConfigurationRoot struct {
	providers []core.ConfigurationProvider
}

func NewConfigurationRoot(providers ...core.ConfigurationProvider) core.ConfigurationRoot {
	root := &ConfigurationRoot{
		providers: providers,
	}
	root.Reload()
	return root
}

func (root *ConfigurationRoot) TryGetConfiguration(providers []core.ConfigurationProvider, k string) (string, bool) {
	ret := ""
	ok := false
	for _, provider := range providers {
		if value, _, ok := provider.TryGet(k); ok {
			ret = value
			ok = true
			break
		}
	}
	return ret, ok
}

func (root *ConfigurationRoot) GetConfiguration(providers []core.ConfigurationProvider, k string) string {
	for _, provider := range providers {
		if value, _, ok := provider.TryGet(k); ok {
			return value
		}
	}
	return ""
}

func (root *ConfigurationRoot) SetConfiguration(providers []core.ConfigurationProvider, k, v string) {
	for _, provider := range providers {
		provider.Set(k, v)
	}
}

func (root *ConfigurationRoot) GetChildrenImplementation(path string) []core.ConfigurationSection {
	results := collections.NewArrayList[string](0)
	for _, provider := range root.providers {
		ks := provider.GetChildKeys(results.ToArray(), path)
		results = collections.NewArrayList[string](len(ks))
		results.AddRange(ks...)
	}
	results = collections.Distinct[string](results, func(a1, a2 any) int {
		if a1 == a2 {
			return 0
		} else {
			return -1
		}
	}).ToList()

	sections := make([]core.ConfigurationSection, results.Size())
	for i, k := range results.ToArray() {
		section := root.GetSection(k)
		if section != nil {
			sections[i] = section
		}
	}
	return sections
}

// ConfigurationRoot Implementations
func (root *ConfigurationRoot) Reload() error {
	for _, provider := range root.providers {
		provider.Load()
	}
	return nil
}
func (root *ConfigurationRoot) Providers() []core.ConfigurationProvider {
	return root.providers
}

// Configuration Implementations
func (root *ConfigurationRoot) GetSection(k string) core.ConfigurationSection {
	return NewConfigurationSection(root, k)
}
func (root *ConfigurationRoot) TryGetString(k string) (string, bool) {
	return root.TryGetConfiguration(root.providers, k)
}
func (root *ConfigurationRoot) GetString(k string) string {
	return root.GetConfiguration(root.providers, k)
}
func (root *ConfigurationRoot) GetChildren() []core.ConfigurationSection {
	return root.GetChildrenImplementation("")
}
