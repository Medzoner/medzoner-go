package app

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/path"
	"github.com/Medzoner/medzoner-go/pkg/infra/services"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/web"
	"github.com/sarulabs/di"
)

// App App
type App struct {
	DebugMode bool
	RootPath  path.RootPath
	Container di.Container
	appWeb    web.IWeb
	database  database.IDbInstance
	conf      config.IConfig
	dbManager database.DbMigration
}

// NewApp is the factory method to create a new App
func NewApp(rootPath path.RootPath, appWeb web.IWeb, database database.IDbInstance, conf config.IConfig, dbManager database.DbMigration) *App {
	return &App{
		RootPath:  rootPath,
		appWeb:    appWeb,
		database:  database,
		conf:      conf,
		dbManager: dbManager,
	}
}

// Handle is the main method to handle the app
func (a *App) Handle(action string) {
	if action == "web" {
		a.appWeb.Start()
	}
	if action == "migrate" {
		a.database.CreateDatabase(a.conf.GetDatabaseName())
		a.dbManager.MigrateUp()
	}
	defer a.deferCT()
	return
}

// LoadContainer LoadContainer
func (a *App) LoadContainer(containerBuilder *di.Builder) {
	err := containerBuilder.Add(services.Service{}.GetDefinitions(a.RootPath)...)
	if err != nil {
		panic(err)
	}
	ct := containerBuilder.Build()
	a.Container = ct
}

func (a *App) deferCT() {
	err := a.Container.Delete()
	if err == nil {
		fmt.Println("ct deleted")
	}
}
