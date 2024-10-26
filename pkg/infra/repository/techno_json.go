package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
)

// TechnoJSONRepository is implementation of TechnoRepository
type TechnoJSONRepository struct {
	Logger   logger.ILogger
	RootPath string
}

// NewTechnoJSONRepository is a constructor
func NewTechnoJSONRepository(logger logger.ILogger, config config.Config) *TechnoJSONRepository {
	return &TechnoJSONRepository{
		Logger:   logger,
		RootPath: string(config.RootPath),
	}
}

// FetchStack FetchStack
func (m *TechnoJSONRepository) FetchStack() (map[string]interface{}, error) {
	jsonFile, err := os.Open(m.RootPath + "pkg/infra/resources/data/jobs/stacks.json")
	if err != nil {
		return nil, fmt.Errorf("error during open json file: %w", err)
	}
	defer m.deferJSONFile(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)
	c := make(map[string]interface{})
	if err = json.Unmarshal(byteValue, &c); err != nil {
		return nil, fmt.Errorf("error during unmarshal json: %w", err)
	}

	return c, nil
}

func (m *TechnoJSONRepository) deferJSONFile(jsonFile *os.File) {
	err := jsonFile.Close()
	if err != nil {
		m.Logger.Error("jsonFile error.")
	}
}
