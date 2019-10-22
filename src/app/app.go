// Copyright 2019 nirajgeorgian. All rights reserved.
// source: src/app/app.go
// package: app

package app

import (
	"github.com/pkg/errors"

	"github.com/nirajgeorgian/account/src/model"

	"github.com/nirajgeorgian/account/src/db"
)

// App is the default app structure
type App struct {
	Config   *Config
	Database *db.Database
}

// New returns new app package instance
func New() (app *App, err error) {
	app = &App{}
	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	dbConfig, err := db.InitConfig()
	if err != nil {
		return nil, err
	}

	app.Database, err = db.New(dbConfig)
	if err != nil {
		return nil, err
	}

	// migrate database on startup
	if err := app.Database.AutoMigrate(&model.AccountORM{}).Error; err != nil {
		return nil, errors.Wrap(err, "unable to automatically migrate migrations table")
	}

	return app, err
}

// Close helps us to close any database connection
func (a *App) Close() error {
	return a.Database.Close()
}
