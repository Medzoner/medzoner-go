package repository

type TechnoRepository interface {
	FetchStack() map[string]interface{}
	FetchExperience() map[string]interface{}
	FetchFormation() map[string]interface{}
	FetchLang() map[string]interface{}
	FetchOther() map[string]interface{}
}
