package restaurantmodel

import (
	"Food-Delivery/common"
	"errors"
	"strings"
)

type RestaurantType string

const TypeNormal RestaurantType = "normal"
const TypePremium RestaurantType = "premium"
const EntityName = "Restaurant"

type Restaurant struct {
	common.SqlModel                    // common.SqlModel `json:",inline"`=> use verison old
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"addr" gorm:"column:addr;"`
	Type            RestaurantType     `json:"type" gorm:"column:type;"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	UserId          int                `json:"-" gorm:"column:user_id;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"` // *common.SimpleUser => dùng để bỏ qua những thông tin của user như email, phone
	LikedCount      int                `json:"liked_count" gorm:"column:liked_count"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

type RestaurantCreate struct {
	common.SqlModel
	Name   string         `json:"name" gorm:"column:name;"` // TODO
	Addr   string         `json:"addr" gorm:"column:addr;"` // address
	UserId int            `json:"-" gorm:"column:user_id;"`
	Logo   *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover  *common.Images `json:"cover" gorm:"column:cover;"`
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

// Trường hợp muốn update 1 filed nào đó về nil ~~ "" thì bắt buộc phải dùng con trỏ *,&
type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name"`
	Addr  *string        `json:"addr" gorm:"column:addr"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName() // table name
}

var (
	ErrNameIsEmpty = errors.New("name cannot be empty")
)
