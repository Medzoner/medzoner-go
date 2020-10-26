package pkg

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"github.com/Medzoner/medzoner-go/pkg/infra/services"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/web"
	"github.com/sarulabs/di"
)

type App struct {
	DebugMode bool
	RootPath  string
	Container di.Container
}

func (a *App) Handle(action string) {
	ct := a.LoadContainer()
	if action == "web" {
		ct.Get("app-web").(web.IWeb).Start()
	}
	if action == "migrate-up" {
		ct.Get("db-manager").(*database.DbMigration).MigrateUp()
	}
	defer a.deferCT(ct)
	return
}

func (a *App) LoadContainer() di.Container {
	builder, _ := di.NewBuilder()
	err := builder.Add(services.Service{}.GetDefinitions(a.RootPath)...)
	if err != nil {
		panic(err)
	}
	ct := builder.Build()
	a.Container = ct
	return ct
}

func (a *App) deferCT(ct di.Container) {
	err := ct.Delete()
	if err == nil {
		fmt.Println("ct deleted")
	}
}
