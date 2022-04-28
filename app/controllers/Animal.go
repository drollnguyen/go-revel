package controllers

import (
	"github.com/revel/revel"
)

type Animal struct {
	*revel.Controller
}

func (c Animal) List() revel.Result {
	return c.Render()
}
