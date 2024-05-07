package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
)

// TechnoJSONRepository TechnoJSONRepository
type TechnoJSONRepository struct {
	Logger   logger.ILogger
	RootPath string
}

// NewTechnoJSONRepository NewTechnoJSONRepository
func NewTechnoJSONRepository(logger logger.ILogger, config config.IConfig) *TechnoJSONRepository {
	return &TechnoJSONRepository{
		Logger:   logger,
		RootPath: string(config.GetRootPath()),
	}
}

// FetchStack FetchStack
func (m *TechnoJSONRepository) FetchStack() map[string]interface{} {
	jsonFile, err := os.Open(m.RootPath + "/pkg/infra/resources/data/jobs/stacks.json")
	if err != nil {
		fmt.Println(err)
	}
	defer m.deferJSONFile(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	c := make(map[string]interface{})
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		panic(err)
	}

	return c
}

// FetchExperience FetchExperience
func (m *TechnoJSONRepository) FetchExperience() map[string]interface{} {
	jsonFile, err := os.Open(m.RootPath + "/pkg/infra/resources/data/jobs/experiences.json")
	if err != nil {
		fmt.Println(err)
	}
	defer m.deferJSONFile(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	c := make(map[string]interface{})
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		panic(err)
	}

	return c
}

// FetchFormation FetchFormation
func (m *TechnoJSONRepository) FetchFormation() map[string]interface{} {
	jsonFile, err := os.Open(m.RootPath + "/pkg/infra/resources/data/jobs/formations.json")
	if err != nil {
		fmt.Println(err)
	}
	defer m.deferJSONFile(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	c := make(map[string]interface{})
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		panic(err)
	}

	return c
}

// FetchLang FetchLang
func (m *TechnoJSONRepository) FetchLang() map[string]interface{} {
	jsonFile, err := os.Open(m.RootPath + "/pkg/infra/resources/data/jobs/langs.json")
	if err != nil {
		fmt.Println(err)
	}
	defer m.deferJSONFile(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	c := make(map[string]interface{})
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		panic(err)
	}

	return c
}

// FetchOther FetchOther
func (m *TechnoJSONRepository) FetchOther() map[string]interface{} {
	jsonFile, err := os.Open(m.RootPath + "/pkg/infra/resources/data/jobs/others.json")
	if err != nil {
		fmt.Println(err)
	}
	defer m.deferJSONFile(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	c := make(map[string]interface{})
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		panic(err)
	}

	return c
}

func (m *TechnoJSONRepository) deferJSONFile(jsonFile *os.File) {
	err := jsonFile.Close()
	if err != nil {
		m.Logger.Error("jsonFile error.")
	}
}
