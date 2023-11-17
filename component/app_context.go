package component

import (
	"bestHabit/component/uploadprovider"

	"github.com/jmoiron/sqlx"
)

type AppContext interface {
	GetMainDBConnection() *sqlx.DB
	SecretKey() string
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db             *sqlx.DB
	secretKey      string
	uploadProvider uploadprovider.UploadProvider
}

func NewAppContext(
	db *sqlx.DB,
	secretKey string,
	uploadProvider uploadprovider.UploadProvider,

) *appCtx {
	return &appCtx{db: db, secretKey: secretKey, uploadProvider: uploadProvider}
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
