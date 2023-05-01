package rstlikebiz

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	restaurantlikemodel "Food-Delivery/module/restaurantlike/model"
	"context"
	"log"
)

// store interface
type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{}, // map[key]value
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type DecLikedCountRestaurantStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

// store biz
type userDislikeRestaurantBiz struct {
	store    UserDislikeRestaurantStore
	decStore DecLikedCountRestaurantStore
}

// new func
func NewUserDislikeRestaurantBiz(
	store UserDislikeRestaurantStore,
	decStore DecLikedCountRestaurantStore,
) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store:    store,
		decStore: decStore,
	}
}

// func call
func (biz *userDislikeRestaurantBiz) DislikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	dataRestaurant, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": data.RestaurantId})

	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if dataRestaurant.LikedCount == 0 {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	if err := biz.store.Delete(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	// TH update bị nghẽn thì nó sẽ chặn API vì vậy phải open grountines
	go func() {
		defer common.AppRecover()
		if err := biz.decStore.DecreaseLikeCount(ctx, data.RestaurantId); err != nil {
			// k chặn đếm like
			log.Println(err)
		}
	}()

	return nil
}
