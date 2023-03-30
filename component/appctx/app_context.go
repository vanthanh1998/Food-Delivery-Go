package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMailDBConnection() *gorm.DB
}

type appCtx struct { // ~~ sqlStore in store.go
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}

func (ctx *appCtx) GetMailDBConnection() *gorm.DB {
	return ctx.db
}
