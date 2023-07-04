package config_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"io/ioutil"
	"os"
	"testing"
)

var conf = config.Config{
	Environment: "test",
	RootPath:    "./../../../",
	DebugMode:   false,
	Options:     nil,
}

func TestConfig(t *testing.T) {
	c := conf
	t.Run("Unit: test config GetRootPath success", func(t *testing.T) {
		c.Init()
		fmt.Println(c.GetRootPath())
		fmt.Println(c.GetEnvironment())
	})
}

func TestEnv(t *testing.T) {
	t.Run("test env found", func(t *testing.T) {
		_, err := ioutil.ReadFile(conf.RootPath + ".env")
		if err != nil {
			fmt.Println(err)
			panic("err")
		}
		err = ioutil.WriteFile(conf.RootPath+".env", []byte([]byte{}), 0644)
		if err != nil {
			fmt.Println(err)
			panic("err")
		}
		conf.Init()
	})
	t.Run("test no env", func(t *testing.T) {
		err := os.Setenv("WAIT_MYSQL", "2")
		if err != nil {
			fmt.Println(err)
			panic("err")
		}
		err = os.Setenv("OPTIONS", "opt")
		if err != nil {
			fmt.Println(err)
			panic("err")
		}
		conf.RootPath = "/tmp"
		conf.Init()
	})
}
