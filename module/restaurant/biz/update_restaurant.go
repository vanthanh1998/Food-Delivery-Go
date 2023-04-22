package restaurantbiz

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
)

// interface
// struct store biz => private
// function new
// hàm chính

type UpdateRestaurant interface {
	Update(ctx context.Context, id int, data restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurant
}

func NewUpdateRestaurantBiz(store UpdateRestaurant) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data restaurantmodel.RestaurantUpdate) error {
	// biz *updateRestaurantBiz dùng để gọi hàm update ở interface

	// validate
	
	/**/
	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
