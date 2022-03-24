package web

type RouteMap struct {
	areas []*RouteArea
}

func NewRouteMap() *RouteMap {
	return &RouteMap{
		areas: make([]*RouteArea, 0),
	}
}

func (r *RouteMap) Area(name string) *RouteArea {
	area := &RouteArea{
		Name:       name,
		RouteItems: make([]*RouteItem, 0),
	}
	r.areas = append(r.areas, area)
	return area
}

type RouteArea struct {
	Name       string
	RouteItems []*RouteItem
}

func (area *RouteArea) Route(method string, path string, handler HTTPHandler) *RouteArea {
	item := &RouteItem{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
	area.RouteItems = append(area.RouteItems, item)
	return area
}

type RouteItem struct {
	Path    string
	Method  string
	Handler HTTPHandler
}
