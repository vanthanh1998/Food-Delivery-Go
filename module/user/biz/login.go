package userbiz

import (
	"Food-Delivery/common"
	"Food-Delivery/component/tokenprovider"
	usermodel "Food-Delivery/module/user/model"
	"context"
)

// interface
type LoginStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	// LoginUser => k nên viết hàm LoginUser ở đây => k dùng đến => vô nghĩa
}

// biz struct => private
type loginBusiness struct {
	storeUser     LoginStore
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

// func new
func NewLoginBusiness(storeUser LoginStore, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. Hash pass form input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)

// func main
func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	passHashed := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayLoad{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	//refreshToken, err := business.tokenProvider.Generate(payload, business.tkCfg.GetRtExp())
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}
	//
	//account := usermodel.NewAccount(accessToken, refreshToken)

	return accessToken, nil
}
