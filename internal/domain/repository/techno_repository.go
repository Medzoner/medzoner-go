//go:generate mockgen -destination=../../../test/mocks/techno_repository.go -package=mocks -source=./techno_repository.go TechnoRepository

package repository

import "context"

type TechnoRepository interface {
	FetchStack(ctx context.Context) (map[string]interface{}, error)
}
