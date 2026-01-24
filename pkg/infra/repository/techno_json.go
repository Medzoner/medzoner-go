package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Medzoner/medzoner-go/internal/config"
)

// TechnoJSONRepository is implementation of TechnoRepository
type TechnoJSONRepository struct {
	RootPath string
}

// NewTechnoJSONRepository is a constructor
func NewTechnoJSONRepository(config config.Config) *TechnoJSONRepository {
	return &TechnoJSONRepository{
		RootPath: string(config.RootPath),
	}
}

// FetchStack FetchStack
func (m *TechnoJSONRepository) FetchStack(ctx context.Context) (map[string]interface{}, error) {
	_ = ctx
	jsonFile, err := os.Open(m.RootPath + "internal/resources/data/jobs/stacks.json")
	if err != nil {
		return nil, fmt.Errorf("error during open json file: %w", err)
	}
	defer m.deferJSONFile(ctx, jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)
	c := make(map[string]interface{})
	if err = json.Unmarshal(byteValue, &c); err != nil {
		return nil, fmt.Errorf("error during unmarshal json: %w", err)
	}

	return c, nil
}

func (m *TechnoJSONRepository) deferJSONFile(ctx context.Context, jsonFile *os.File) {
	if err := jsonFile.Close(); err != nil {
		//m.Telemetry.Error(ctx, "jsonFile error.")
	}
}
