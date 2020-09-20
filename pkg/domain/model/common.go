package model

type ICommon interface {
	GetId() int
	SetId(id int) ICommon
	GetUuid() string
	SetUuid(id string) ICommon
}
