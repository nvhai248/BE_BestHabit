package component

import (
	"bestHabit/component/uploadprovider"
	"bestHabit/pubsub"

	"github.com/jmoiron/sqlx"
)

type AppContext interface {
	GetMainDBConnection() *sqlx.DB
	SecretKey() string
	UploadProvider() uploadprovider.UploadProvider
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db             *sqlx.DB
	secretKey      string
	uploadProvider uploadprovider.UploadProvider
	pb             pubsub.Pubsub
}

func NewAppContext(
	db *sqlx.DB,
	secretKey string,
	uploadProvider uploadprovider.UploadProvider,
	pb pubsub.Pubsub,

) *appCtx {
	return &appCtx{db: db, secretKey: secretKey, uploadProvider: uploadProvider, pb: pb}
}

func (ctx *appCtx) GetMainDBConnection() *sqlx.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) GetPubSub() pubsub.Pubsub {
	return ctx.pb
}
