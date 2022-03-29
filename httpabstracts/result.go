package httpabstracts

type HttpResult interface {
	Write(ctx *HttpContext)
}
