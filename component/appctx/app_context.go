package appctx

import (
	"Food-Delivery/component/uploadprovider"
	"Food-Delivery/pubsub"
	"Food-Delivery/skio"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMailDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
	GetRealTimeEngine() skio.RealTimeEngine
}

type appCtx struct { // ~~ sqlStore in storage.go
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
	rtEngine       skio.RealTimeEngine
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
func (ctx *appCtx) GetRealTimeEngine() skio.RealTimeEngine {
	return ctx.rtEngine
}

func (ctx *appCtx) SetRealTimeEngine(rt skio.RealTimeEngine) {
	ctx.rtEngine = rt
}
