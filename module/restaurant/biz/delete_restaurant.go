package restaurantbiz

import (
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
	"errors"
)

// interface của golang thông thường khai báo ở nơi mà chúng ta dùng nó
type DeleteRestaurantStore interface { // có thể dùng N interface này
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{}, // map[key]value
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error // storage/create.go
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore // khai báo store interface viết ở trên
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz { // *deleteRestaurantBiz: return con trỏ *deleteRestaurantBiz
	return &deleteRestaurantBiz{store: store} // bắt buộc phải có &
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	// get old data
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}

	if err := biz.store.Delete(context, id); err != nil {
		return err
	}

	return nil
}
