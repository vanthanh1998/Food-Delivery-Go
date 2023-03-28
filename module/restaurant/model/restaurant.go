package restaurantmodel

type Restaurant struct {
	Id     int    `json:"id" gorm:"column:id"`         // id of the restaurant
	Name   string `json:"name" gorm:"column:name"`     // TODO
	Addr   string `json:"addr" gorm:"column:addr"`     // address
	Status int    `json:"status" gorm:"column:status"` // status
}

func (Restaurant) TableName() string {
	return "restaurants" // table name
}

type RestaurantCreate struct {
	Id   int    `json:"id" gorm:"column:id"`     // id of the restaurant
	Name string `json:"name" gorm:"column:name"` // TODO
	Addr string `json:"addr" gorm:"column:addr"` // address
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName() // table name
}

// Trường hợp muốn update 1 filed nào đó về nil ~~ "" thì bắt buộc phải dùng con trỏ *,&
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName() // table name
}
