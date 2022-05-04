package controllers

import (
	"github.com/revel/revel"
)

type Animal struct {
	*revel.Controller
}

func (c Animal) Index() revel.Result {
	return c.Render()
}

func (c Animal) Create() revel.Result {
	return c.Render()
}
