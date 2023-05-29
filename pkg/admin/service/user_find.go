package service

import (
	"context"
	"ecommerce/internal/config/database"
	"ecommerce/internal/util/mgquery"
	"ecommerce/pkg/admin/dao"
	responsemodel "ecommerce/pkg/admin/model/response"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
)

// All ...
func (s userImpl) All(ctx context.Context, q mgquery.AppQuery) (res responsemodel.UserAll) {
	var (
		d   = dao.D()
		wg  = sync.WaitGroup{}
		col = database.GetUserCol()
	)

	cond := bson.M{}
	q.AssignKeyword(cond)

	wg.Add(1)
	go func() {
		defer wg.Done()

		res.List = make([]responsemodel.UserBrief, 0)

		findOpts := q.GetFindOptionsWithPage()

		docs := d.FindByCondition(ctx, col, cond, findOpts)

		res.List = s.getListUserBrief(ctx, docs)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res.Total = d.CountByCondition(ctx, col, cond)
	}()

	wg.Wait()

	res.Limit = q.Limit
	return
}
