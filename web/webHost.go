package web

import (
	"net/http"
	"servicelocator"
)

type WebHost struct {
	middleware MiddlewareDelegate
	routeMap   *RouteMap
	server     http.Server
	rootScope  servicelocator.ServiceScope
}
