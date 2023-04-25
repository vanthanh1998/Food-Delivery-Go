package restaurantmodel

type Filter struct {
	OwnerId int `json:"user_id,omitempty" form:"user_id"` // form => use querystring trên đường dẫn
}
