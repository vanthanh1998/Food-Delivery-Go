package memcache

import (
	usermodel "Food-Delivery/module/user/model"
	"context"
	"fmt"
	"log"
)

// interface store
type RealStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

// struct

type userCaching struct { // chính là interface trong file caching.go
	store     Caching
	realStore RealStore
}

// func new
func NewUserCaching(store Caching, realStore RealStore) *userCaching {
	return &userCaching{
		store:     store,
		realStore: realStore,
	}
}

// func main

func (uc *userCaching) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	userId := conditions["id"].(int)
	key := fmt.Sprintf("user-%d", userId)
	log.Println("-------------")
	log.Println(key)
	log.Println("-------------")
	userInCache := uc.store.Read(key)

	if userInCache != nil {
		return userInCache.(*usermodel.User), nil
	}
	
	user, err := uc.realStore.FindUser(ctx, conditions, moreInfo...)

	if err != nil {
		return nil, err
	}

	// update cache
	uc.store.Write(key, user)

	return user, nil
}
