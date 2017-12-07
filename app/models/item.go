package models

import (
	"github.com/jinzhu/gorm"
)

// Item model
type Item struct {
	gorm.Model
	Name  string
	About string
	Image string
}
