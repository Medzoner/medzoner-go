package entity

import "github.com/Medzoner/medzoner-go/pkg/domain/model"

type CommonModel struct {
	id int `json:"id" db:"id"`
}

func (c *CommonModel) Id() int {
	return c.id
}

func (c *CommonModel) SetId(id int) model.ICommon {
	c.id = id
	return c
}
