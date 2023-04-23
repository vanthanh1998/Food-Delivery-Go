package userbiz

import (
	"Food-Delivery/common"
	usermodel "Food-Delivery/module/user/model"
	"context"
)

// interface
type RegisterStore interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

// interface hash pw
type Hasher interface {
	Hash(data string) string
}

// biz struct => thường dùng để gọi các interface đc khai báo trong file của nó thui => private
type registerBusiness struct {
	registerStore RegisterStore
	hasher        Hasher
}

// new function:
// 1. Biến bên trong hàm là những store struct đc khai báo bên trong {}
// 2. Luôn return ha store struct
func NewRegisterBusiness(registerStore RegisterStore, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStore: registerStore, // dùng để call func trong folder store
		hasher:        hasher,
	}
}

// ham call bên tầng transport
func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	// user, _ ~~ user, error => dấu _ được sd để bỏ qua error
	user, _ := business.registerStore.FindUser(ctx, map[string]interface{}{"email": data.Email})

	// if user == nil => return false
	// if user != nil => return true
	if user != nil {
		//if user.Status == 0 {
		//	return errors.New("error user has been disable")
		//}
		return usermodel.ErrEmailExisted
	}
	salt := common.GenSalt(50)
	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code
	//data.Status = 1

	if err := business.registerStore.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
