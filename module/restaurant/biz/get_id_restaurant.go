package restaurantbiz

import (
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
)

type GetIdRestaurantStore interface {
	GetId(ctx context.Context, id int, data *restaurantmodel.Restaurant) error // interface gọi thẳng hàm thực hiện chính của tầng biz
}

type getIdRestaurantBiz struct {
	store GetIdRestaurantStore // struct {}
}

func NewGetIdRestaurantBiz(store GetIdRestaurantStore) *getIdRestaurantBiz {
	return &getIdRestaurantBiz{store: store} // hàm để trả về
}

func (biz *getIdRestaurantBiz) GetIdRestaurant(ctx context.Context, id int, data *restaurantmodel.Restaurant) error {
	if err := biz.store.GetId(ctx, id, data); err != nil {
		return err
	}
	return nil
}
