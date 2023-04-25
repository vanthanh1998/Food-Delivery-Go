package restaurantbiz

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
)

// interface của golang thông thường khai báo ở nơi mà chúng ta dùng nó
type DeleteRestaurantStore interface { // có thể dùng N interface này
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{}, // map[key]value
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error // storage/create.go
}

type deleteRestaurantBiz struct {
	store     DeleteRestaurantStore // khai báo store interface viết ở trên
	requester common.Requester
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore, requester common.Requester) *deleteRestaurantBiz { // *deleteRestaurantBiz: return con trỏ *deleteRestaurantBiz
	return &deleteRestaurantBiz{store: store, requester: requester} // bắt buộc phải có &
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	// get old data
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil { // TH1
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 { // TH2
		// because TH1 return err => return new
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if oldData.UserId != biz.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, nil)
	}

	return nil
}
