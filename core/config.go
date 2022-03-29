package core

import "strings"

type ConfigurationBuilder interface {
	Properties() map[string]any
	Sources() []ConfigurationSource
	AddSource(source ConfigurationSource)
	Build() (ConfigurationRoot, error)
}

type ConfigurationSource interface {
	Build(builder ConfigurationBuilder) ConfigurationProvider
}

type Configuration interface {
	TryGetString(k string) (string, bool)
	GetString(k string) string
	GetSection(k string) ConfigurationSection
	GetChildren() []ConfigurationSection
}

type ConfigurationSection interface {
	Configuration

	Key() string
	Value() string
	Path() string
	HasValue() bool
}

type ConfigurationRoot interface {
	Configuration

	Reload() error
	Providers() []ConfigurationProvider
}

var DefaultConfigPathKeyDelimiter = ":"

func CombineConfigPath(pathSegments ...string) string {
	if len(pathSegments) == 0 {
		return ""
	}
	return strings.Join(pathSegments, DefaultConfigPathKeyDelimiter)
}

func GetSectionKeyFromConfigPath(path string) string {
	if path == "" {
		return ""
	}
	lastDelimiterIndex := strings.LastIndex(path, DefaultConfigPathKeyDelimiter)
	if lastDelimiterIndex == -1 {
		return path
	}
	return path[lastDelimiterIndex+1:]
}

func GetParentConfigPath(path string) string {
	if path == "" {
		return ""
	}
	lastDelimiterIndex := strings.LastIndex(path, DefaultConfigPathKeyDelimiter)
	if lastDelimiterIndex == -1 {
		return ""
	}
	return path[:lastDelimiterIndex]
}

type ConfigurationProvider interface {
	TryGet(k string) (string, ConfigurationProvider, bool)
	Set(k string, v string)
	Load()
	GetChildKeys(earlierKeys []string, parentPath string) []string
}
