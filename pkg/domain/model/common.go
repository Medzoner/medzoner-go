package model

type ICommon interface {
	GetID() int
	SetID(id int) ICommon
	GetUUID() string
	SetUUID(id string) ICommon
}
