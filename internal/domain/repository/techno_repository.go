package repository

import "context"

// TechnoRepository TechnoRepository
//
//go:generate mockgen -destination=../../../test/mocks/techno_repository.go -package=domainRepositoryMock -source=./techno_repository.go TechnoRepository
type TechnoRepository interface {
	FetchStack(ctx context.Context) (map[string]interface{}, error)
}
