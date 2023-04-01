package restaurantstorage

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
) (*restaurantmodel.Restaurant, error) { // if isData return pointer data else isNotData return error
	// why có con trỏ ???
	// Giả sử nếu k dùng con trỏ thì sẽ k return về nil được mà bắt buộc phải return về structure rỗng
	// structure rỗng: default value int: 0, string:"", blool:"false" => dẫn đến mất memory
	// Trường hợp nữa là: if ở tham số đầu tiên =! nil => restaurantmodel.Restaurant nó chỉ là structure rỗng => tốn memory
	//  => Nếu k phải là pointer(con trỏ) thì nó sẽ có default value
	var data restaurantmodel.Restaurant // data

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		// return restaurantmodel.Restaurant{}, err
		// trường hợp First thì cần phải if ra 2 trường hợp để check cho rõ
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
