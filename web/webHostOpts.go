package web

import (
	"github.com/ah-its-andy/downton/core"
	"github.com/ah-its-andy/downton/hosting"
)

type WebHostOptions struct {
	primaryConfig  core.Configuration
	fallbackConfig core.Configuration
	environment    hosting.HostEnv

	ApplicationName string
	ContentRootPath string
	Environment     string
	ShutdownTimeout int
}
