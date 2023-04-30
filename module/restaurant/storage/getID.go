package restaurantstorage

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
)

func (s *sqlStore) GetId(ctx context.Context, id int, data *restaurantmodel.Restaurant, moreKeys ...string) error {

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
