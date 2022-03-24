package web

type MiddlewareDelegate func(ctx *HttpContext, next MiddlewareDelegate)

type Middleware interface {
	Handle(ctx *HttpContext, next MiddlewareDelegate)
}
