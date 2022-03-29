package config

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/ah-its-andy/downton/core"
	"gopkg.in/yaml.v2"
)

type YamlConfigurationProvider struct {
	BinaryConfigurationProvider
}

func NewYamlConfigurationProvider(source *BinaryConfigurationSource) *YamlConfigurationProvider {
	if source.buildFunc == nil {
		source.buildFunc = buildYamlProvider
	}
	provider := &YamlConfigurationProvider{
		BinaryConfigurationProvider: *NewBinaryConfigurationProvider(
			source, loadYamlConfig),
	}
	return provider
}

func buildYamlProvider(source *BinaryConfigurationSource, builder core.ConfigurationBuilder) core.ConfigurationProvider {
	return NewYamlConfigurationProvider(source)
}

func loadYamlConfig(provider *BinaryConfigurationProvider, reader io.Reader) {
	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(fmt.Sprintf("Failed to read config file: %s", err))
	}
	values := make(map[string]interface{})
	err = yaml.Unmarshal(buffer, &values)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse config file with yaml format: %s", err))
	}
	for key, value := range values {
		provider.SetValue("", key, value)
	}
}
