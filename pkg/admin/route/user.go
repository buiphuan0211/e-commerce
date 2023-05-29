package route

import (
	"ecommerce/pkg/admin/handler"
	routevalidation "ecommerce/pkg/admin/route/validation"
	"github.com/labstack/echo/v4"
)

// user ...
func user(e *echo.Group) {
	var (
		g = e.Group("/users")
		h = handler.User()
		v = routevalidation.User()
	)

	// Create
	g.POST("", h.Create, v.Create)

	// All
	g.GET("", h.All, v.All)

}
