package rstlikebiz

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	restaurantlikemodel "Food-Delivery/module/restaurantlike/model"
	"golang.org/x/net/context"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{}, // map[key]value
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type IncLikedCountRestaurantStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncLikedCountRestaurantStore
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	incStore IncLikedCountRestaurantStore,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:    store,
		incStore: incStore,
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

	// TH update bị nghẽn thì nó sẽ chặn API vì vậy phải open grountines
	go func() {
		defer common.AppRecover()
		if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
			// k chặn đếm like
			log.Println(err)
		}
	}()

	return nil
}
