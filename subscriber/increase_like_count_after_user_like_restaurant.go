package subscriber

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	restaurantstorage "Food-Delivery/module/restaurant/storage"
	"context"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	//GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)

	db := appCtx.GetMailDBConnection()

	store := restaurantstorage.NewSQLStore(db)

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c // rút data vào msg
			likeData := msg.Data().(HasRestaurantId)
			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}
