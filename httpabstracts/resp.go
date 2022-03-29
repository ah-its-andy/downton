package httpabstracts

type HttpResponse interface {
	StatusCode() int
	SetStatusCode(int)
	Body() []byte
	SetBody([]byte)
	Headers() map[string]string
	SetHeader(string, string)
	SetHeaders(map[string]string)
	ContentType() string
	SetContentType(string)
}
