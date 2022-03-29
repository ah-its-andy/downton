package config

import "github.com/ah-its-andy/downton/core"

type ConfigurationBuilder struct {
	properties map[string]any
	sources    []core.ConfigurationSource
}

func (builder ConfigurationBuilder) Properties() map[string]any {
	return builder.properties
}
func (builder ConfigurationBuilder) Sources() []core.ConfigurationSource {
	ret := make([]core.ConfigurationSource, len(builder.sources))
	copy(ret, builder.sources)
	return ret
}
func (builder ConfigurationBuilder) AddSource(source core.ConfigurationSource) {
	builder.sources = append(builder.sources, source)
}
func (builder ConfigurationBuilder) Build() (core.ConfigurationRoot, error) {
	providers := make([]core.ConfigurationProvider, len(builder.sources))
	for i, source := range builder.sources {
		providers[i] = source.Build(builder)
	}
	return NewConfigurationRoot(providers...), nil
}
