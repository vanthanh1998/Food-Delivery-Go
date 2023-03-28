package restaurantstorage

import "gorm.io/gorm"

type sqlStore struct { // k viết hoa chữ cái đầu để k public ra ngoài
	db *gorm.DB // database
}

func NewSQLStore(db *gorm.DB) *sqlStore { // public ra ngoài
	return &sqlStore{db: db}
}
