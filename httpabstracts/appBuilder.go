package httpabstracts

type AppBuilder interface {
	UseMiddleware(middleware func(RequestFunc, RequestFunc) error)
	Items() map[string]interface{}
	Build() RequestFunc
}
