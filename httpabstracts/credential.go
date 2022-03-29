package httpabstracts

type AuthenticationCredential interface {
	Verify(ctx *HttpContext) bool
	Get(k string) string
	Set(k string, v string)
}
