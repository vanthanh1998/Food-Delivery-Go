package userstore

import (
	"Food-Delivery/common"
	usermodel "Food-Delivery/module/user/model"
	"context"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	// db.Begin() // open transaction => TH sd thực hiện nhiều lệnh xuống table users
	// db.RollBack() => bắt buộc phải có rollback
	// db.Commit // => save db
	// db.RollBack() => bắt buộc phải có rollback
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
