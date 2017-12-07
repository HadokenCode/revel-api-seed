package controllers

import (
	"net/http"

	"github.com/obitux/revel-api-seed/app/models"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type Item struct {
	gormc.TxnController
}

func init() {
	revel.InterceptFunc(checkUser, revel.BEFORE, &Item{})
}

// List all Items
func (c Item) List() revel.Result {
	var items = []models.Item{}
	c.Txn.Find(&items)
	return c.RenderJSON(items)
}

// Create a Item
func (c Item) Create() revel.Result {
	type Data struct {
		Name  string
		About string
		Image string
	}
	jsonData := Data{}
	c.Params.BindJSON(&jsonData)
	item := models.Item{Name: jsonData.Name, About: jsonData.About, Image: jsonData.Image}
	c.Txn.Create(&item)
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(item)
}

// Read a Item
func (c Item) Read(id int) revel.Result {
	item := models.Item{}
	c.Txn.First(&item, id)
	if item.ID == 0 {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON("not found")
	}
	return c.RenderJSON(item)
}

// Update a Item
func (c Item) Update(id int) revel.Result {
	// TODO
	return c.RenderJSON("not implemented yet")
}

// Delete a Item
func (c Item) Delete(id int) revel.Result {
	item := models.Item{}
	c.Txn.First(&item, id)
	if item.ID == 0 {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON("not found")
	}
	c.Txn.Delete(&item)
	return c.RenderJSON("deleted")
}
