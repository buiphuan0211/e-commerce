package dao

import (
	"context"
	mgmodel "ecommerce/internal/model/mg"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DInterface ...
type DInterface interface {
	Create(ctx context.Context, mgCol *mongo.Collection, payload interface{}) error
	FindOnByCondition(ctx context.Context, mgCol *mongo.Collection, cond interface{}, opts ...*options.FindOneOptions) (doc mgmodel.User, err error)
	FindByID(ctx context.Context, mgCol *mongo.Collection, id primitive.ObjectID) (doc *mgmodel.User, err error)
	FindByCondition(ctx context.Context, mgCol *mongo.Collection, cond interface{}, opts ...*options.FindOptions) (docs []mgmodel.User)
	CountByCondition(ctx context.Context, mgCol *mongo.Collection, cond interface{}) (total int64)
	UpdateOneByCondition(ctx context.Context, mgCol *mongo.Collection, cond interface{}, payload interface{}) (err error)
}

// dImpl ...
type dImpl struct{}

// D ...
func D() DInterface {
	return dImpl{}
}
