package dao

import (
	"context"
	"ecommerce/internal/config/database"
	"ecommerce/internal/config/plogger"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateOneByCondition ...
func (dImpl) UpdateOneByCondition(ctx context.Context, mgCol *mongo.Collection, cond interface{}, payload interface{}) (err error) {
	var col = database.GetUserCol()
	if _, err = col.UpdateByID(ctx, cond, payload); err != nil {
		contentError := fmt.Sprintf("dao.%s - UpdateOneByCondition", mgCol.Name())
		plogger.Error(contentError, plogger.LogData{
			Message: err.Error(),
			Data: plogger.Map{
				"cond":    cond,
				"payload": payload,
			},
		})
	}

	return
}
