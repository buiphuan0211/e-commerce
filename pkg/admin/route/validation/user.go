package routevalidation

import (
	"ecommerce/internal/response"
	"ecommerce/internal/util/echocontext"
	requestmodel "ecommerce/pkg/admin/model/request"
	"github.com/labstack/echo/v4"
)

// UserInterface ...
type UserInterface interface {
	Create(next echo.HandlerFunc) echo.HandlerFunc
	All(next echo.HandlerFunc) echo.HandlerFunc
}

// userImpl ...
type userImpl struct{}

// User ...
func User() UserInterface {
	return userImpl{}
}

// Create ...
func (userImpl) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.UserCreate

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetPayload(c, payload)
		return next(c)
	}
}

// All ...
func (userImpl) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query requestmodel.UserAll
		if err := c.Bind(&query); err != nil {
			return response.R400(c, nil, "")
		}

		if err := query.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetQuery(c, query)
		return next(c)
	}
}
