package web

import (
	"github.com/ah-its-andy/downton/config"
	"github.com/ah-its-andy/downton/hosting"
)

type WebHostOptions struct {
	primaryConfig  config.Configuration
	fallbackConfig config.Configuration
	environment    hosting.HostEnv

	ApplicationName string
	ContentRootPath string
	Environment     string
	ShutdownTimeout int
}
