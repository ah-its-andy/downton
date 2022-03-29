package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/ah-its-andy/downton/core"
)

type ConfigurationBuilder struct {
	properties map[string]any
	sources    []core.ConfigurationSource
}

func (builder *ConfigurationBuilder) Properties() map[string]any {
	return builder.properties
}
func (builder *ConfigurationBuilder) Sources() []core.ConfigurationSource {
	ret := make([]core.ConfigurationSource, len(builder.sources))
	copy(ret, builder.sources)
	return ret
}
func (builder *ConfigurationBuilder) AddSource(source core.ConfigurationSource) {
	builder.sources = append(builder.sources, source)
}
func (builder *ConfigurationBuilder) Build() (core.ConfigurationRoot, error) {
	providers := make([]core.ConfigurationProvider, len(builder.sources))
	for i, source := range builder.sources {
		providers[i] = source.Build(builder)
	}
	return NewConfigurationRoot(providers...), nil
}

// extensions
func (builder *ConfigurationBuilder) AddMemorySource(initialData map[string]string) {
	if initialData == nil {
		builder.AddSource(NewMemoryConfigurationSource(initialData))
	} else {
		builder.AddSource(NewMemoryConfigurationSource(make(map[string]string)))
	}
}

func (builder *ConfigurationBuilder) AddCommandLine(args []string, switchMapping map[string]string) {
	builder.AddSource(NewCmdLineConfigurationSource(switchMapping, args))
}

func (builder *ConfigurationBuilder) AddYamlFile(fileProvider core.FileProvider) {
	builder.AddSource(NewYamlConfigurationSource(fileProvider))
}

func (builder *ConfigurationBuilder) AddJSONFile(fileProvider core.FileProvider) {
	builder.AddSource(NewJSONConfigurationSource(fileProvider))
}

func (builder *ConfigurationBuilder) AddDir(path string, recursive bool) {
	entries, err := os.ReadDir(path)
	if err != nil {
		panic(fmt.Sprintf("Failed to read directory %s: %s", path, err))
	}
	for _, entry := range entries {
		if entry.IsDir() && !recursive {
			continue
		}

		if entry.IsDir() {
			builder.AddDir(entry.Name(), recursive)
		} else {
			file, err := entry.Info()
			if err != nil {
				panic(fmt.Sprintf("Failed to read file %s: %s", entry.Name(), err))
			}
			if strings.HasSuffix(file.Name(), ".yaml") || strings.HasSuffix(file.Name(), ".yml") {
				builder.AddYamlFile(core.NewLocalFileProvider(file.Name()))
			} else if strings.HasSuffix(file.Name(), ".json") {
				builder.AddJSONFile(core.NewLocalFileProvider(file.Name()))
			} else {
				panic(fmt.Sprintf("Unknown file extension for %s", file.Name()))
			}
		}
	}
}
