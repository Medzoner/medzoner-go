package model

type ICommon interface {
	Id() int
	SetId(id int) ICommon
}
