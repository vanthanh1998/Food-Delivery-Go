package appctx

import (
	"Food-Delivery/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMailDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appCtx struct { // ~~ sqlStore in store.go
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider, secretKey string) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider, secretKey: secretKey}
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
