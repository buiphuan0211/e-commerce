package route

import (
	"ecommerce/internal/util/routemiddleware"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.Use(routemiddleware.CORSConfig())

	//e.Use(routeauth.Jwt())

	r := e.Group("/admin/ecommerce")
	common(r)
	user(r)

}
