package restaurantmodel

import (
	"Food-Delivery/common"
	"errors"
	"strings"
)

type RestaurantType string

const TypeNormal RestaurantType = "normal"
const TypePremium RestaurantType = "premium"

type Restaurant struct {
	common.SqlModel
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
	Type string `json:"type" gorm:"column:type;"`
}

func (Restaurant) TableName() string {
	return "restaurants" // table name
}

type RestaurantCreate struct {
	common.SqlModel
	Name string `json:"name" gorm:"column:name;"` // TODO
	Addr string `json:"addr" gorm:"column:addr;"` // address
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName() // table name
}

// Trường hợp muốn update 1 filed nào đó về nil ~~ "" thì bắt buộc phải dùng con trỏ *,&
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName() // table name
}

var (
	ErrNameIsEmpty = errors.New("name cannot be empty")
)
