package subscriber

import (
	"Food-Delivery/component/appctx"
	restaurantstorage "Food-Delivery/module/restaurant/storage"
	"Food-Delivery/pubsub"
	"context"
)

//func DecreaseLikeCountAfterUserDislikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserDislikeRestaurant)
//
//	db := appCtx.GetMailDBConnection()
//
//	store := restaurantstorage.NewSQLStore(db)
//
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c // rút data vào msg
//			likeData := msg.Data().(HasRestaurantId)
//			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
//		}
//	}()
//}

func DecreaseLikeCountAfterUserDislikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user dislike restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			db := appCtx.GetMailDBConnection()
			store := restaurantstorage.NewSQLStore(db)
			likeData := message.Data().(HasRestaurantId)

			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
