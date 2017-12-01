package controllers

import (
	"goseed/app/models"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type Node struct {
	gormc.TxnController
}

func init() {
	revel.InterceptFunc(checkUser, revel.BEFORE, &Node{})
}

func (c Node) List() revel.Result {
	var nodes = []models.Node{}
	c.Txn.Find(&nodes)
	return c.RenderJSON(nodes)
}

func (c Node) Create() revel.Result {
	return c.RenderJSON("not implemented yet")
}
