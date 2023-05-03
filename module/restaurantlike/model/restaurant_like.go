package restaurantlikemodel

import (
	"Food-Delivery/common"
	"fmt"
	"time"
)

type RestaurantType string

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at,omitempty" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false;"` // *common.SimpleUser => dùng để bỏ qua những thông tin của user như email, phone
}

// table name
func (Like) TableName() string { return "restaurant_likes" }

// get restaurant id
func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

// get user id
func (l *Like) GetUserId() int {
	return l.UserId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrCannotDislikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot dislike this restaurant"),
		fmt.Sprintf("ErrCannotDislikeRestaurant"),
	)
}
