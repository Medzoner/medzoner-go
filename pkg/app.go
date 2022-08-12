package pkg

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/services"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/web"
	"github.com/sarulabs/di"
)

//App App
type App struct {
	DebugMode  bool
	RootPath   string
	Container  di.Container
	DbInstance database.IDbInstance
	AppWeb     web.IWeb
}

//Handle Handle
func (a *App) Handle(action string) {
	if action == "web" {
		a.Container.Get("app-web").(web.IWeb).Start()
	}
	if action == "migrate" {
		a.DbInstance.CreateDatabase(
			a.Container.Get("config").(config.IConfig).GetDatabaseName(),
		)
		a.Container.Get("db-manager").(*database.DbMigration).MigrateUp()
	}
	defer a.deferCT()
	return
}

//LoadContainer LoadContainer
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
