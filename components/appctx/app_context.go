package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
}

type appContext struct {
	db     *gorm.DB
	secret string
}

func NewAppContext(db *gorm.DB, secret string) *appContext {
	return &appContext{db: db, secret: secret}
}

func (appCtx *appContext) GetMainDBConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) SecretKey() string {
	return appCtx.secret
}
