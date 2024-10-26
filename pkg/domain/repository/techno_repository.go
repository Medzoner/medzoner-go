package repository

// TechnoRepository TechnoRepository
//
//go:generate mockgen -destination=../../../test/mocks/pkg/domain/repository/techno_repository.go -package=domainRepositoryMock -source=./techno_repository.go TechnoRepository
type TechnoRepository interface {
	FetchStack() (map[string]interface{}, error)
}
