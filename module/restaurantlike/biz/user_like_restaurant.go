package rstlikebiz

import (
	restaurantlikemodel "Food-Delivery/module/restaurantlike/model"
	"golang.org/x/net/context"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
	}
}

func (biz *userLikeRestaurantBiz) UserLikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	return nil
}
