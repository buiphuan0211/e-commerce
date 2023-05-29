package service

import (
	"context"
	"ecommerce/internal/config/database"
	"ecommerce/pkg/admin/dao"
	requestmodel "ecommerce/pkg/admin/model/request"
	responsemodel "ecommerce/pkg/admin/model/response"
)

// Create ...
func (s userImpl) Create(ctx context.Context, payload requestmodel.UserCreate) (res responsemodel.Upsert, err error) {

	var (
		d   = dao.D()
		doc = payload.ConvertToBSON()
		col = database.GetUserCol()
	)

	if err = d.Create(ctx, col, doc); err != nil {
		return
	}

	res.ID = doc.ID.Hex()
	return
}
