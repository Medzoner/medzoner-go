package config_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	var conf = config.Config{
		Environment: "test",
		RootPath:    "./../../../",
		DebugMode:   false,
		Options:     nil,
	}
	t.Run("test env found", func(t *testing.T) {
		_, err := os.ReadFile(string(conf.RootPath + ".env"))
		if err != nil {
			fmt.Println(err)
			panic("err")
		}
		err = os.WriteFile(string(conf.RootPath+".env"), []byte([]byte{}), 0644)
		if err != nil {
			fmt.Println(err)
			panic("err")
		}
		_, err = conf.Init()
		if err != nil {
			t.Error(err)
		}
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
		_, err = conf.Init()
		if err != nil {
			t.Error(err)
		}
	})
}
