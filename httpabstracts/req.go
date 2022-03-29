package httpabstracts

type HttpRequest interface {
	Scheme() string
	Method() string
	Path() string
	Host() string
	Headers() map[string]string
	Body() []byte
	Query() map[string]string
	Form() map[string]string
	IsForm() bool
	ContentType() string
}
