package dao

import (
	"context"
	"ecommerce/internal/config/plogger"
	mgmodel "ecommerce/internal/model/mg"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOnByCondition ...
func (dImpl) FindOnByCondition(ctx context.Context, mgCol *mongo.Collection, cond interface{}, opts ...*options.FindOneOptions) (doc mgmodel.User, err error) {
	if err = mgCol.FindOne(ctx, cond, opts...).Decode(&doc); err != nil {
		contentError := fmt.Sprintf("dao.%s - FindOnByCondition", mgCol.Name())
		plogger.Error(contentError, plogger.LogData{
			Message: err.Error(),
			Data: plogger.Map{
				"cond": cond,
			},
		})
	}

	return
}

// FindByID ...
func (dImpl) FindByID(ctx context.Context, mgCol *mongo.Collection, id primitive.ObjectID) (doc *mgmodel.User, err error) {
	if err = mgCol.FindOne(ctx, bson.M{"_id": id}).Decode(&doc); err != nil {
		contentError := fmt.Sprintf("dao.%s - FindByID", mgCol.Name())
		plogger.Error(contentError, plogger.LogData{
			Message: err.Error(),
			Data: plogger.Map{
				"id": id,
			},
		})
	}

	return
}

// FindByCondition ...
func (dImpl) FindByCondition(ctx context.Context, mgCol *mongo.Collection, cond interface{}, opts ...*options.FindOptions) (docs []mgmodel.User) {
	cursor, err := mgCol.Find(ctx, cond, opts...)
	if err != nil {
		contentError := fmt.Sprintf("dao.%s - FindByCondition(col.Find)", mgCol.Name())
		plogger.Error(contentError, plogger.LogData{
			Message: err.Error(),
			Data: plogger.Map{
				"cond": cond,
				"opts": opts,
			},
		})
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &docs); err != nil {
		contentError := fmt.Sprintf("dao.%s - FindByCondition(cursor.All)", mgCol.Name())
		plogger.Error(contentError, plogger.LogData{
			Message: err.Error(),
			Data: plogger.Map{
				"cond": cond,
				"opts": opts,
			},
		})
	}

	return
}

// CountByCondition ...
func (dImpl) CountByCondition(ctx context.Context, mgCol *mongo.Collection, cond interface{}) (total int64) {
	total, _ = mgCol.CountDocuments(ctx, cond)
	return
}
