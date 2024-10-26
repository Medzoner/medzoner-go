package config_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	var conf, err = config.NewConfig()
	if err != nil {
		fmt.Println(err)
		panic("err")
	}
	conf.Environment = "test"
	conf.RootPath = "./../../../"
	conf.DebugMode = false
	conf.Options = nil

	t.Run("test env found", func(t *testing.T) {
		_, err := os.ReadFile(string(conf.RootPath + ".env"))
		if err != nil {
			t.Errorf("expected %v, got %v", nil, err)
		}
		err = os.WriteFile(string(conf.RootPath+".env"), []byte([]byte{}), 0644)
		if err != nil {
			t.Errorf("expected %v, got %v", nil, err)
		}
	})
}
