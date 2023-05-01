package subscriber

import (
	"Food-Delivery/component/appctx"
	"context"
)

func Setup(appCtx appctx.AppContext, ctx context.Context) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, ctx)
	//DecreaseLikeCountAfterUserLikeRestaurant(appCtx, ctx)
}
