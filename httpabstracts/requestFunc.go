package httpabstracts

type RequestFunc func(ctx *HttpContext, next RequestFunc)
