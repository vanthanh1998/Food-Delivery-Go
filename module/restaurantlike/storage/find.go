package restaurantlikestorage

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{}, // map[key]value
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	// find nhà hàng đó đã được like hay chưa
	var data restaurantmodel.Restaurant
	db := s.db

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
