package restaurantlikestorage

import (
	"Food-Delivery/common"
	restaurantlikemodel "Food-Delivery/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, userId int, restaurantId int) error {
	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
