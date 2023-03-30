package restaurantbiz

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
)

// interface của golang thông thường khai báo ở nơi mà chúng ta dùng nó
type ListRestaurantStore interface { // có thể dùng N interface này
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error) // storage/list.go
}

type listRestaurantBiz struct {
	store ListRestaurantStore // khai báo store interface viết ở trên
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz { // *deleteRestaurantBiz: return con trỏ *deleteRestaurantBiz
	return &listRestaurantBiz{store: store} // bắt buộc phải có &
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) { // có 2 tham số []restaurantmodel.Restaurant, error) thì return về như phía dưới
	result, err := biz.store.ListDataWithCondition(context, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
