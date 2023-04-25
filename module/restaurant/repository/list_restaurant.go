package restaurantrepo

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
	"log"
)

type ListRestaurantStore interface { // có thể dùng N interface này
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error) // storage/list.go
}

type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
	// GetRestaurantLikesOld(ctx context.Context, ids []int) ([]restaurantlikemodel.Like, error)
	// => bắt buộc dùng vòng lặp for ra 2 vòng mới tìm được ID nhà hàng có bao nhiêu lượt like
}

type listRestaurantRepo struct {
	store     ListRestaurantStore // khai báo store interface viết ở trên
	likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore LikeRestaurantStore) *listRestaurantRepo { // *deleteRestaurantBiz: return con trỏ *deleteRestaurantBiz
	return &listRestaurantRepo{store: store, likeStore: likeStore} // bắt buộc phải có &
}

func (biz *listRestaurantRepo) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) { // có 2 tham số []restaurantmodel.Restaurant, error) thì return về như phía dưới
	result, err := biz.store.ListDataWithCondition(context, filter, paging, "User")

	if err != nil {
		return nil, err
	}

	// xử lý hàm GetRestaurantLikes => mai xem :))
	ids := make([]int, len(result))

	for i := range ids {
		ids[i] = result[i].Id
	}

	likeMap, err := biz.likeStore.GetRestaurantLikes(context, ids)

	// nếu có err thì có thể bỏ qua lỗi này và continue
	if err != nil {
		log.Println(err) // if is error => like_count return = 0
		return nil, err
	}

	for i, item := range result {
		result[i].LikeCount = likeMap[item.Id]
	}

	return result, nil
}
