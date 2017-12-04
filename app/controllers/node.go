package controllers

import (
	"net/http"

	"github.com/obitux/revel-api-seed/app/models"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type Node struct {
	gormc.TxnController
}

func init() {
	revel.InterceptFunc(checkUser, revel.BEFORE, &Node{})
}

// List all Nodes
func (c Node) List() revel.Result {
	var nodes = []models.Node{}
	c.Txn.Find(&nodes)
	return c.RenderJSON(nodes)
}

// Create a Node
func (c Node) Create() revel.Result {
	type Data struct {
		Name string
		IP   string
	}
	jsonData := Data{}
	c.Params.BindJSON(&jsonData)
	node := models.Node{Name: jsonData.Name, IP: jsonData.IP}
	c.Txn.Create(&node)
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(node)
}

// Read a Node
func (c Node) Read(id int) revel.Result {
	node := models.Node{}
	c.Txn.First(&node, id)
	if node.ID == 0 {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON("not found")
	}
	return c.RenderJSON(node)
}

// Update a Node
func (c Node) Update(id int) revel.Result {
	// TODO
	return c.RenderJSON("not implemented yet")
}

// Delete a Node
func (c Node) Delete(id int) revel.Result {
	node := models.Node{}
	c.Txn.First(&node, id)
	if node.ID == 0 {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON("not found")
	}
	c.Txn.Delete(&node)
	return c.RenderJSON("deleted")
}
