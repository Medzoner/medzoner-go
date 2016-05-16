package repository

import (
	"encoding/json"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"os"
)

type TechnoJsonRepository struct {
	Conn     *sqlx.DB
	Logger   logger.ILogger
	RootPath string
}

func (m *TechnoJsonRepository) FetchStack() map[string]interface{} {
	jsonFile, err := os.Open(m.RootPath + "/pkg/infra/resources/data/jobs/stacks.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	c := make(map[string]interface{})
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		panic(err)
	}

	return c
}

func (m *TechnoJsonRepository) FetchExperience() map[string]interface{} {
	jsonFile, err := os.Open(m.RootPath + "/pkg/infra/resources/data/jobs/experiences.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	c := make(map[string]interface{})
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		panic(err)
	}

	return c
}

func (m *TechnoJsonRepository) FetchFormation() map[string]interface{} {
	jsonFile, err := os.Open(m.RootPath + "/pkg/infra/resources/data/jobs/formations.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	c := make(map[string]interface{})
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		panic(err)
	}

	return c
}

func (m *TechnoJsonRepository) FetchLang() map[string]interface{} {
	jsonFile, err := os.Open(m.RootPath + "/pkg/infra/resources/data/jobs/langs.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	c := make(map[string]interface{})
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		panic(err)
	}

	return c
}

func (m *TechnoJsonRepository) FetchOther() map[string]interface{} {
	jsonFile, err := os.Open(m.RootPath + "/pkg/infra/resources/data/jobs/others.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	c := make(map[string]interface{})
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		panic(err)
	}

	return c
}
