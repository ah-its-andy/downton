package config

import "github.com/ah-its-andy/downton/core"

type ConfigurationSection struct {
	root core.ConfigurationRoot
	path string
	k    string
}

func NewConfigurationSection(root core.ConfigurationRoot, path string) core.ConfigurationSection {
	return &ConfigurationSection{
		root: root,
		path: path,
		k:    core.GetSectionKeyFromConfigPath(path),
	}
}

func (section *ConfigurationSection) Key() string {
	return section.k
}
func (section *ConfigurationSection) Value() string {
	return section.root.GetString(section.k)
}
func (section *ConfigurationSection) Path() string {
	return section.path
}
func (section *ConfigurationSection) HasValue() bool {
	_, ok := section.root.TryGetString(section.k)
	return ok
}

// Configuration Implementations
func (section *ConfigurationSection) TryGetString(k string) (string, bool) {
	return section.root.TryGetString(core.CombineConfigPath(section.path, k))
}
func (section *ConfigurationSection) GetString(k string) string {
	return section.root.GetString(core.CombineConfigPath(section.path, k))
}
func (section *ConfigurationSection) GetSection(k string) core.ConfigurationSection {
	return section.root.GetSection(core.CombineConfigPath(section.path, k))
}
func (section *ConfigurationSection) GetChildren() []core.ConfigurationSection {
	return section.root.GetChildren()
}
