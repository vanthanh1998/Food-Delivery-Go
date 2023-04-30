package rstlikebiz

import (
	restaurantlikemodel "Food-Delivery/module/restaurantlike/model"
	"context"
)

// store interface
type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}

// store biz
type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
}

// new func
func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store}
}

// func call
func (biz *userDislikeRestaurantBiz) DislikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := biz.store.Delete(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	return nil
}
