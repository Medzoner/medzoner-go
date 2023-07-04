package model

// ICommon ICommon
type ICommon interface {
	GetID() int
	SetID(id int) ICommon
	GetUUID() string
	SetUUID(id string) ICommon
}
