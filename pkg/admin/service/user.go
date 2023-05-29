package service

import (
	"context"
	"ecommerce/internal/util/mgquery"
	requestmodel "ecommerce/pkg/admin/model/request"
	responsemodel "ecommerce/pkg/admin/model/response"
)

// UserInterface ...
type UserInterface interface {
	Create(ctx context.Context, payload requestmodel.UserCreate) (res responsemodel.Upsert, err error)
	All(ctx context.Context, q mgquery.AppQuery) (res responsemodel.UserAll)
}

// userImpl ...
type userImpl struct {
}

// User ...
func User() UserInterface {
	return userImpl{}
}
