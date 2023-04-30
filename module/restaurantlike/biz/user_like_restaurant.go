package rstlikebiz

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	restaurantlikemodel "Food-Delivery/module/restaurantlike/model"
	"golang.org/x/net/context"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{}, // map[key]value
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
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

	dataRestaurant, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": data.RestaurantId})

	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if dataRestaurant.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	return nil
}
