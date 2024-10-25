package repository

import (
	"testing"
)

func TestTechnoJSONRepository(t *testing.T) {
	t.Run("Unit: test FetchStack failed", func(t *testing.T) {
		repo := &TechnoJSONRepository{
			RootPath: "failed",
		}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		repo.FetchStack()
	})
}
