package appctx

import (
	"Food-Delivery/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMailDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct { // ~~ sqlStore in store.go
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider}
}

func (ctx *appCtx) GetMailDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}
