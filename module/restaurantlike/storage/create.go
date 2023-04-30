package restaurantlikestorage

import (
	"Food-Delivery/common"
	restaurantlikemodel "Food-Delivery/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db
	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	//db.Exec => tăng lượt like khi create restaurant_like
	//=> cách này sai =>  Trong storage của restaurant_like chỉ thao tác vs DB restaurant_like => k update restaurant
	//db.Exec("Update restaurants SET liked_count = liked_count + 1 WHERE id = ?", data.RestaurantId)

	return nil
}
