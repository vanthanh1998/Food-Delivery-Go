package subscriber

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	restaurantstorage "Food-Delivery/module/restaurant/storage"
	"context"
)

func DecreaseLikeCountAfterUserDislikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserDislikeRestaurant)

	db := appCtx.GetMailDBConnection()

	store := restaurantstorage.NewSQLStore(db)

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c // rút data vào msg
			likeData := msg.Data().(HasRestaurantId)
			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}
