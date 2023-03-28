package restaurantbiz

import (
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
	"errors"
)

// interface của golang thông thường khai báo ở nơi mà chúng ta dùng nó
type CreateRestaurantStore interface { // có thể dùng N interface này
	Create(context context.Context, data *restaurantmodel.RestaurantCreate) error // storage/create.go
}

type createRestaurantBiz struct {
	store CreateRestaurantStore // khai báo store interface viết ở trên
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz { // *deleteRestaurantBiz: return con trỏ *deleteRestaurantBiz
	return &createRestaurantBiz{store: store} // bắt buộc phải có &
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New(" Name cannot be empty")
	}
	if err := biz.store.Create(context, data); err != nil {
		return err
	}

	return nil
}
