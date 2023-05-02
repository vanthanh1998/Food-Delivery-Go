package subscriber

import (
	"Food-Delivery/component/appctx"
	restaurantstorage "Food-Delivery/module/restaurant/storage"
	"Food-Delivery/pubsub"
	"context"
	"log"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	//GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			db := appCtx.GetMailDBConnection()
			store := restaurantstorage.NewSQLStore(db)
			likeData := message.Data().(HasRestaurantId)

			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func PushNotificationWhenUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Push notification when user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			log.Println("Push notification when user like restaurant id: ", likeData.GetRestaurantId())

			return nil
		},
	}
}
