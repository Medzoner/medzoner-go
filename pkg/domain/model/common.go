package model

type ICommon interface {
	Id() int
	SetId(id int) ICommon
	Uuid() string
	SetUuid(id string) ICommon
}
