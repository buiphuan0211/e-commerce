package service

import (
	"context"
	mgmodel "ecommerce/internal/model/mg"
	"ecommerce/internal/util/ptime"
	responsemodel "ecommerce/pkg/admin/model/response"
	"sync"
)

// getListCategoryBrief ...
func (s userImpl) getListUserBrief(ctx context.Context, docs []mgmodel.User) (res []responsemodel.UserBrief) {
	var (
		wg    = sync.WaitGroup{}
		total = len(docs)
	)
	res = make([]responsemodel.UserBrief, total)

	wg.Add(total)
	for i, doc := range docs {
		go func(index int, user mgmodel.User) {
			defer wg.Done()
			res[index] = s.brief(ctx, user)
		}(i, doc)
	}

	wg.Wait()
	return
}

// brief ...
func (userImpl) brief(ctx context.Context, user mgmodel.User) responsemodel.UserBrief {
	return responsemodel.UserBrief{
		ID:        user.ID.Hex(),
		Name:      user.Name,
		CreatedAt: ptime.TimeResponseInit(user.CreatedAt),
		UpdatedAt: ptime.TimeResponseInit(user.UpdatedAt),
	}
}
