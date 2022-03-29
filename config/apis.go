package config

import "github.com/ah-its-andy/downton/core"

func AddMemorySource(builder core.ConfigurationBuilder, initialData map[string]string) {
	if initialData == nil {
		builder.AddSource(NewMemoryConfigurationSource(initialData))
	}
}
