package app

import (
	"context"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/path"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/web"
)

// App App
type App struct {
	DebugMode bool
	RootPath  path.RootPath
	appWeb    web.IWeb
	Database  database.IDbInstance
	conf      config.IConfig
	DbManager database.DbMigration
}

// NewApp is the factory method to create a new App
func NewApp(rootPath path.RootPath, appWeb web.IWeb, database database.IDbInstance, conf config.IConfig, dbManager database.DbMigration) *App {
	return &App{
		RootPath:  rootPath,
		appWeb:    appWeb,
		Database:  database,
		conf:      conf,
		DbManager: dbManager,
	}
}

// Handle is the main method to handle the app
func (a *App) Handle(action string) {
	if action == "web" {
		a.appWeb.Start()
	}
	if action == "migrate" {
		a.Database.CreateDatabase(a.conf.GetDatabaseName())
		a.DbManager.MigrateUp()
	}
	defer a.deferCT()
	return
}

// StopServer StopServer
func (a *App) StopServer(ctx context.Context) error {
	return a.appWeb.StopServer(ctx)
}

func (a *App) deferCT() {
	fmt.Println("ct deleted")
}
