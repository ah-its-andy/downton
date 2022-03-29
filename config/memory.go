package config

import (
	"fmt"

	"github.com/ah-its-andy/downton/core"
)

type MemoryConfigurationSource struct {
	initialData map[string]string
}

func NewMemoryConfigurationSource(initialData map[string]string) *MemoryConfigurationSource {
	if initialData == nil {
		return &MemoryConfigurationSource{
			initialData: make(map[string]string),
		}
	} else {
		return &MemoryConfigurationSource{
			initialData: initialData,
		}
	}
}

func (source *MemoryConfigurationSource) Build(builder core.ConfigurationBuilder) core.ConfigurationProvider {
	return NewMemoryConfigurationProvider(source)
}

type MemoryConfigurationProvider struct {
	ConfigurationProvider

	source *MemoryConfigurationSource
}

func NewMemoryConfigurationProvider(source *MemoryConfigurationSource) *MemoryConfigurationProvider {
	if source == nil {
		panic(fmt.Sprintf("NewMemoryConfigurationProvider called with nil source"))
	}
	ret := &MemoryConfigurationProvider{
		source:                source,
		ConfigurationProvider: *NewConfigurationProvider(),
	}
	if len(source.initialData) == 0 {
		ret.source.initialData = make(map[string]string)
		for k, v := range source.initialData {
			ret.data[k] = v
		}
	}
	return ret
}

func (provider *MemoryConfigurationProvider) Add(k, v string) {
	provider.data[k] = v
}
