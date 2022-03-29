package config

import (
	"os"
	"strings"

	"github.com/ah-its-andy/downton/core"
)

type EnvironmentConfigurationSource struct {
	prefixs []string
}

func (source *EnvironmentConfigurationSource) Build(builder core.ConfigurationBuilder) core.ConfigurationProvider {
	return NewEnvironmentConfigurationProvider(source)
}

type EnvironmentConfigurationProvider struct {
	ConfigurationProvider

	source           *EnvironmentConfigurationSource
	normalizePrefixs []string
}

func NewEnvironmentConfigurationProvider(source *EnvironmentConfigurationSource) *EnvironmentConfigurationProvider {
	ret := &EnvironmentConfigurationProvider{
		source:                source,
		normalizePrefixs:      make([]string, len(source.prefixs)),
		ConfigurationProvider: *NewConfigurationProvider(),
	}
	for i, prefix := range source.prefixs {
		ret.normalizePrefixs[i] = ret.normalize(prefix)
	}
	return ret
}

func (provider *EnvironmentConfigurationProvider) Load() {
	for _, env := range os.Environ() {
		separatorIndex := strings.Index(env, "=")
		k, v := "", ""
		if separatorIndex < 0 {
			k = env
		} else {
			k = env[separatorIndex:]
			v = env[separatorIndex+1:]
		}
		if provider.includeWithPrefixs(provider.normalize(k)) {
			provider.Set(provider.normalize(k), v)
		}
	}
}

func (provider *EnvironmentConfigurationProvider) includeWithPrefixs(k string) bool {
	for _, prefix := range provider.normalizePrefixs {
		if strings.HasPrefix(k, prefix) {
			return true
		}
	}
	return false
}

func (provider *EnvironmentConfigurationProvider) normalize(k string) string {
	return strings.Replace(k, "__", core.DefaultConfigPathKeyDelimiter, -1)
}
