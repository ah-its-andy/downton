package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/ah-its-andy/downton/core"
)

func NewJSONConfigurationSource(fileProvider core.FileProvider) *BinaryConfigurationSource {
	return NewBinaryConfigurationSource(fileProvider, buildJSONConfigurationProvider)
}

type JSONConfigurationProvider struct {
	BinaryConfigurationProvider
}

func NewJSONConfigurationProvider(source *BinaryConfigurationSource) *JSONConfigurationProvider {
	if source.buildFunc == nil {
		source.buildFunc = buildYamlProvider
	}
	provider := &JSONConfigurationProvider{
		BinaryConfigurationProvider: *NewBinaryConfigurationProvider(
			source, loadYamlConfig),
	}
	return provider
}

func buildJSONConfigurationProvider(source *BinaryConfigurationSource, builder core.ConfigurationBuilder) core.ConfigurationProvider {
	return NewYamlConfigurationProvider(source)
}

func loadJSONConfiguration(provider *BinaryConfigurationProvider, reader io.Reader) {
	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(fmt.Sprintf("Failed to read config file: %s", err))
	}
	values := make(map[string]interface{})
	err = json.Unmarshal(buffer, &values)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse config file with json format: %s", err))
	}
	for key, value := range values {
		provider.SetValue("", key, value)
	}
}
