package restaurantstorage

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
)

// step 1: Fetch 1000 ids => caching
// step 2: fetch 50 first ids => data (page 1)
// step 3: Page 2, omit 50 ids and fetch the next 50 ids

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)") // get status active = 1

	// check conditions
	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("user_id = ?", f.OwnerId)
		}
	}

	// count trước rồi mới paging
	if err := db.Count(&paging.Total).Error; err != nil { // truyền con trỏ vào để nó có thể update đc total
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FaceCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
