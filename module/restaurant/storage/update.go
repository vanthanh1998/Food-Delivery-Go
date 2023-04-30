package restaurantstorage

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) Update(ctx context.Context, id int, data restaurantmodel.RestaurantUpdate) error {
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	// update ++ GORM
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	// update -- GORM
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
