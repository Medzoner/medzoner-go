package repository

// TechnoRepository TechnoRepository
type TechnoRepository interface {
	FetchStack() map[string]interface{}
}
