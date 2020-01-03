package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteRegister struct {
	Path    string
	Handler http.Handler
}

type RouteServiceProvider struct {
	Router   *mux.Router
	Register []RouteRegister
}

func (routeServiceProvider *RouteServiceProvider) SetRouter(router *mux.Router) {
	routeServiceProvider.Router = router
}

// GetRouterInstance get instance of gorila mux route
func (routeServiceProvider RouteServiceProvider) GetRouterInstance() *mux.Router {
	return routeServiceProvider.Router
}

func (routeServiceProvider *RouteServiceProvider) RegisterRoute(path string, handler http.Handler) {
	routeServiceProvider.Register = append(routeServiceProvider.Register, RouteRegister{
		Path:    path,
		Handler: handler,
	})
}

func (routeServiceProvider RouteServiceProvider) InitRoute() {
	for _, element := range routeServiceProvider.Register {
		fmt.Println("init route: " + element.Path)
		var route = routeServiceProvider.GetRouterInstance()
		route.Handle(element.Path, element.Handler)
	}
}
