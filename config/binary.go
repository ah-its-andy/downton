package config

import (
	"fmt"
	"io"
	"strings"

	"github.com/ah-its-andy/downton/core"
)

type BinaryConfigurationSource struct {
	fileProvider core.FileProvider

	buildFunc func(*BinaryConfigurationSource, core.ConfigurationBuilder) core.ConfigurationProvider
}

func NewBinaryConfigurationSource(fileProvider core.FileProvider,
	buildFunc func(*BinaryConfigurationSource, core.ConfigurationBuilder) core.ConfigurationProvider) *BinaryConfigurationSource {
	return &BinaryConfigurationSource{
		fileProvider: fileProvider,
		buildFunc:    buildFunc,
	}
}

func (source *BinaryConfigurationSource) Build(builder core.ConfigurationBuilder) core.ConfigurationProvider {
	return source.buildFunc(source, builder)
}

type BinaryConfigurationProvider struct {
	ConfigurationProvider

	source       *BinaryConfigurationSource
	binaryLoader func(provider *BinaryConfigurationProvider, reader io.Reader)

	loaded bool
}

func NewBinaryConfigurationProvider(source *BinaryConfigurationSource, binaryLoader func(provider *BinaryConfigurationProvider, reader io.Reader)) *BinaryConfigurationProvider {
	provider := &BinaryConfigurationProvider{
		source:       source,
		binaryLoader: binaryLoader,
	}
	return provider
}

func (provider *BinaryConfigurationProvider) Load() {
	if provider.loaded {
		panic(fmt.Sprintf("BinaryConfigurationProvider.Load() called more than once"))
	}

	if provider.source == nil {
		panic(fmt.Sprintf("BinaryConfigurationProvider.Load() called with nil source"))
	}

	if provider.source.fileProvider == nil {
		panic(fmt.Sprintf("BinaryConfigurationProvider.Load() called with nil fileProvider"))
	}

	if provider.binaryLoader == nil {
		panic(fmt.Sprintf("BinaryConfigurationProvider.Load() called with nil binaryLoader"))
	}

	reader, err := provider.source.fileProvider.Reader()
	if err != nil {
		panic(fmt.Sprintf("Failed to open config file: %s", err))
	}

	provider.binaryLoader(provider, reader)
	provider.loaded = true
}

func (provider *BinaryConfigurationProvider) SetValue(path string, k string, v any) {
	paths := make([]string, 0)
	if path != "" {
		paths = append(paths, path)
	}
	paths = append(paths, k)
	nPath := strings.Join(paths, core.DefaultConfigPathKeyDelimiter)
	if mapV, ok := v.(map[string]any); ok {
		for key, value := range mapV {
			provider.SetValue(nPath, key, value)
		}
	} else if sliceV, ok := v.([]any); ok {
		for i, value := range sliceV {
			provider.SetValue(nPath, fmt.Sprintf("$%d", i), value)
		}
	} else {
		provider.Set(nPath, fmt.Sprintf("%v", v))
	}
}
