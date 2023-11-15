package component

import (
	"github.com/jmoiron/sqlx"
)

type AppContext interface {
	GetMainDBConnection() *sqlx.DB
	SecretKey() string
}

type appCtx struct {
	db        *sqlx.DB
	secretKey string
}

func NewAppContext(
	db *sqlx.DB,
	secretKey string,

) *appCtx {
	return &appCtx{db: db, secretKey: secretKey}
}

func (ctx *appCtx) GetMainDBConnection() *sqlx.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}
