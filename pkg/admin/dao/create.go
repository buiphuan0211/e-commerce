package dao

import (
	"context"
	"ecommerce/internal/config/plogger"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create ...
func (dImpl) Create(ctx context.Context, mgCol *mongo.Collection, payload interface{}) error {
	if _, err := mgCol.InsertOne(ctx, payload); err != nil {
		contentError := fmt.Sprintf("dao.%s - InsertOne", mgCol.Name())
		plogger.Error(contentError, plogger.LogData{
			Source:  "ecommerce",
			Message: err.Error(),
			Data: plogger.Map{
				"payload": payload,
			},
		})
	}

	return nil
}
