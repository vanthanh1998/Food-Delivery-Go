package restaurantbiz

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
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
	result, err := biz.repo.ListRestaurant(context, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
