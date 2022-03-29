package httpabstracts

import (
	"context"

	servicelocator "github.com/ah-its-andy/downton/serviceLocator"
)

type HttpContext struct {
	Request  HttpRequest
	Response HttpResponse

	ClientInfo *ClientInfo
	Credential AuthenticationCredential

	Items map[string]any

	CancelationContext context.Context

	ServiceScope servicelocator.ServiceScope
}

func (ctx *HttpContext) Cancel() error {
	return nil
}
