package context

import (
	"context"
	"gorm.io/gorm"
)

type AppContext struct {
	context.Context
	db *gorm.DB
}

func NewAppContext(ctx context.Context) *AppContext {
	return &AppContext{
		Context: ctx,
	}
}

func SetDB(ctx context.Context, db *gorm.DB) {
	appContext, ok := ctx.(*AppContext)
	if !ok {
		return
	}
	appContext.db = db
}

func GetDB(ctx context.Context) *gorm.DB {
	appContext, ok := ctx.(*AppContext)
	if !ok {
		return nil
	}
	return appContext.db
}
