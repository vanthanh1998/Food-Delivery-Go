package restaurantstorage

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
)

// hàm CreateRestaurant chỉ là method của struct (s *sqlStore) => nên khi gọi sqlStore thì hàm CreateRestaurant sẽ auto run
func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
