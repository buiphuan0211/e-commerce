package handler

import (
	"ecommerce/internal/response"
	"ecommerce/internal/util/echocontext"
	"ecommerce/internal/util/mgquery"
	requestmodel "ecommerce/pkg/admin/model/request"
	"ecommerce/pkg/admin/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// UserInterface ...
type UserInterface interface {
	Create(c echo.Context) error
	All(c echo.Context) error
}

// userImpl ...
type userImpl struct{}

// User ...
func User() UserInterface {
	return userImpl{}
}

// Create doc
// @tags User
// @summary Create
// @id admin-user-create
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body requestmodel.UserCreate true "Payload"
// @success 200 {object} responsemodel.Upsert
// @router /users [post]
func (userImpl) Create(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.UserCreate)
		s       = service.User()
	)

	res, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}

// All godoc
// @tags User
// @summary All
// @id admin-user-all
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload query requestmodel.UserAll true "Query"
// @success 200 {object} responsemodel.UserAll
// @router /users [GET]
func (userImpl) All(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		qParams = echocontext.GetQuery(c).(requestmodel.UserAll)
		s       = service.User()
	)

	q := mgquery.AppQuery{
		Page:          qParams.Page,
		Limit:         qParams.Limit,
		Keyword:       qParams.Keyword,
		SortInterface: bson.M{"createdAt": -1},
	}

	res := s.All(ctx, q)
	return response.R200(c, res, "")
}
