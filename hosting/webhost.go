package hosting

import (
	"github.com/ah-its-andy/downton/core"
)

type WebHost interface {
	ServiceScope() core.ServiceScope

	Start()
	Stop()
}
