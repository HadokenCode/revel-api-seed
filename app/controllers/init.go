package controllers

import (
	"github.com/obitux/revel-api-seed/app/models"

	gorm "github.com/revel/modules/orm/gorm/app"
	"github.com/revel/revel"
)

func initializeDB() {
	gorm.DB.AutoMigrate(&models.User{})
	var firstUser = models.User{Name: "Demo", Email: "demo@demo.com", Active: true}
	firstUser.SetNewPassword("demo")
	gorm.DB.Create(&firstUser)

	gorm.DB.AutoMigrate(&models.Item{})
}

func init() {
	revel.OnAppStart(initializeDB)
}
