package restaurantstorage

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
)

func (s *sqlStore) GetId(ctx context.Context, id int, data *restaurantmodel.Restaurant) error {
	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
