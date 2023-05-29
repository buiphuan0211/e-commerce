package server

import (
	"ecommerce/internal/config/plogger"
	"ecommerce/pkg/admin/route"
	"ecommerce/pkg/admin/server/initialize"
	"github.com/labstack/echo/v4"
)

func Boostrap(e *echo.Echo) {
	// Init logger ...
	plogger.Init("e-ecommerce", "ecommerce-admin")

	// Init modules
	initialize.Init()

	// Routes
	route.Init(e)
}
