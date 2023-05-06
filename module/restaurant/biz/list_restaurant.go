package restaurantbiz

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
	"go.opencensus.io/trace"
)

type ListRestaurantRepo interface { // có thể dùng N interface này
	ListRestaurant(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo // khai báo storage interface viết ở trên
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) { // có 2 tham số []restaurantmodel.Restaurant, error) thì return về như phía dưới
	ctx1, span := trace.StartSpan(context, "biz.list_restaurant")

	span.AddAttributes(
		trace.Int64Attribute("page", int64(paging.Page)),
		trace.Int64Attribute("limit", int64(paging.Limit)),
	)

	result, err := biz.repo.ListRestaurant(ctx1, filter, paging)

	span.End()

	if err != nil {
		return nil, err
	}

	return result, nil
}
