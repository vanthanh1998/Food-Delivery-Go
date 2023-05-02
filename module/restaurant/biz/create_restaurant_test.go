package restaurantbiz

import (
	"Food-Delivery/common"
	restaurantmodel "Food-Delivery/module/restaurant/model"
	"context"
	"errors"
	"testing"
)

type mockCreateRestaurant struct{}

func (mockCreateRestaurant) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "Thanh" {
		return common.ErrDB(errors.New("something went wrong in DB"))
	}

	data.Id = 200
	return nil
}

func TestNewCreateRestaurantBiz(t *testing.T) {
	biz := NewCreateRestaurantBiz(mockCreateRestaurant{})

	dataTest := restaurantmodel.RestaurantCreate{Name: ""}
	err := biz.CreateRestaurant(context.Background(), &dataTest)

	if err == nil || err.Error() != "invalid request" {
		t.Errorf("Failed")
	}

	dataTest = restaurantmodel.RestaurantCreate{Name: "Thanh"}
	err = biz.CreateRestaurant(context.Background(), &dataTest)

	if err == nil {
		t.Errorf("Failed")
	}

	dataTest = restaurantmodel.RestaurantCreate{Name: "200Lab"}
	err = biz.CreateRestaurant(context.Background(), &dataTest)

	if err != nil {
		t.Errorf("Failed")
	}
}
