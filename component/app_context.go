package component

import (
	"bestHabit/component/cronjob"
	"bestHabit/component/mailprovider"
	"bestHabit/component/oauthprovider"
	"bestHabit/component/uploadprovider"
	"bestHabit/pubsub"

	"github.com/jmoiron/sqlx"
)

type AppContext interface {
	GetMainDBConnection() *sqlx.DB
	SecretKey() string
	UploadProvider() uploadprovider.UploadProvider
	GetPubSub() pubsub.Pubsub
	GetEmailSender() mailprovider.EmailSender
	GetGGOAuth() oauthprovider.GGOAuthProvider
	GetCronJob() cronjob.CronJobProvider
}

type appCtx struct {
	db             *sqlx.DB
	secretKey      string
	uploadProvider uploadprovider.UploadProvider
	pb             pubsub.Pubsub
	emailSender    mailprovider.EmailSender
	oauthProvider  oauthprovider.GGOAuthProvider
	cronProvider   cronjob.CronJobProvider
}

func NewAppContext(
	db *sqlx.DB,
	secretKey string,
	uploadProvider uploadprovider.UploadProvider,
	pb pubsub.Pubsub,
	emailSender mailprovider.EmailSender,
	oauthProvider oauthprovider.GGOAuthProvider,
	cronProvider cronjob.CronJobProvider,
) *appCtx {
	return &appCtx{db: db,
		secretKey:      secretKey,
		uploadProvider: uploadProvider,
		pb:             pb, emailSender: emailSender,
		oauthProvider: oauthProvider,
		cronProvider:  cronProvider}
}

func (ctx *appCtx) GetCronJob() cronjob.CronJobProvider {
	return ctx.cronProvider
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

func (ctx *appCtx) GetEmailSender() mailprovider.EmailSender {
	return ctx.emailSender
}

func (ctx *appCtx) GetGGOAuth() oauthprovider.GGOAuthProvider {
	return ctx.oauthProvider
}
