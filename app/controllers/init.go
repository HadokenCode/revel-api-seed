package controllers

import (
	"goseed/app/models"

	gorm "github.com/revel/modules/orm/gorm/app"
	"github.com/revel/revel"
)

func initializeDB() {
	gorm.DB.AutoMigrate(&models.User{})
	var firstUser = models.User{Name: "Demo", Email: "demo@demo.com"}
	firstUser.SetNewPassword("demo")
	firstUser.Active = true
	gorm.DB.Create(&firstUser)

	gorm.DB.AutoMigrate(&models.Node{})
}

func init() {
	revel.OnAppStart(initializeDB)
}
