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
	t.Run("Unit: test FetchExperience failed", func(t *testing.T) {
		repo := &TechnoJSONRepository{
			RootPath: "failed",
		}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		repo.FetchExperience()
	})
	t.Run("Unit: test FetchFormation failed", func(t *testing.T) {
		repo := &TechnoJSONRepository{
			RootPath: "failed",
		}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		repo.FetchFormation()
	})
	t.Run("Unit: test FetchLang failed", func(t *testing.T) {
		repo := &TechnoJSONRepository{
			RootPath: "failed",
		}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		repo.FetchLang()
	})
	t.Run("Unit: test FetchOther failed", func(t *testing.T) {
		repo := &TechnoJSONRepository{
			RootPath: "failed",
		}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		repo.FetchOther()
	})
}
