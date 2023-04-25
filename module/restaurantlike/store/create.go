package restaurantlikestore

import (
	"Food-Delivery/common"
	restaurantlikemodel "Food-Delivery/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data restaurantlikemodel.Like) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
