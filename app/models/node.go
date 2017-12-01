package models

import (
	"github.com/jinzhu/gorm"
)

// Node model
type Node struct {
	gorm.Model
	Name   string `gorm:"size:255"`
	IP     string `gorm:"size:42"`
	Active bool
}
