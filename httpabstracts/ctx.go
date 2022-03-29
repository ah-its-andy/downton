package httpabstracts

import (
	"context"

	"github.com/ah-its-andy/downton/core"
)

type HttpContext struct {
	Request  HttpRequest
	Response HttpResponse

	ClientInfo *ClientInfo
	Credential AuthenticationCredential

	Items map[string]any

	CancelationContext context.Context

	ServiceScope core.ServiceScope
}

func (ctx *HttpContext) Cancel() error {
	return nil
}
