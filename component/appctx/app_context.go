package appctx

import (
	"Food-Delivery/component/uploadprovider"
	"Food-Delivery/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMailDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appCtx struct { // ~~ sqlStore in storage.go
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
}

func NewAppContext(
	db *gorm.DB, uploadProvider uploadprovider.UploadProvider,
	secretKey string,
	ps pubsub.Pubsub,
) *appCtx {
	return &appCtx{
		db:             db,
		uploadProvider: uploadProvider,
		secretKey:      secretKey,
		ps:             ps,
	}
}

func (ctx *appCtx) GetMailDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetPubSub() pubsub.Pubsub {
	return ctx.ps
}
