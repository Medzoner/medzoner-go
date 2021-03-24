package pkg

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/services"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/web"
	"github.com/sarulabs/di"
)

//App App
type App struct {
	DebugMode bool
	RootPath  string
	Container di.Container
}

//Handle Handle
func (a *App) Handle(action string) {
	if action == "web" {
		a.Container.Get("app-web").(web.IWeb).Start()
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
