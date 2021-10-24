package routes

import (
	"spos/stores/controllers"

	"github.com/labstack/echo"
	"github.com/s-pos/go-utils/middleware"
)

type route struct {
	middleware middleware.Clients
	controller controllers.Controller
}

type Route interface {
	Router() *echo.Echo
}

func NewRoute(mdl middleware.Clients, ctrl controllers.Controller) Route {
	return &route{
		middleware: mdl,
		controller: ctrl,
	}
}

func (r *route) Router() *echo.Echo {
	router := echo.New()

	store := router.Group("", echo.WrapMiddleware(r.middleware.Logger), echo.WrapMiddleware(r.middleware.Session))
	store.POST("/", r.controller.NewStoreHandler)

	return router
}
