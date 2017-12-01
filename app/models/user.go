package models

import (
	"fmt"
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	gorm.Model
	Name           string `gorm:"size:255"`
	Email          string `gorm:"type:varchar(100);unique_index"`
	HashedPassword []byte
	Active         bool
	FileName       string `gorm:"size:255"`
}

// SetNewPassword set a new hashsed password to user
func (user *User) SetNewPassword(passwordString string) {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	user.HashedPassword = bcryptPassword
}

func (user *User) String() string {
	return fmt.Sprintf("User(%s)", user.Email)
}

var emailRegexp = regexp.MustCompile(`[a-zA-Z0-9_\-]+@[a-zA-Z0-9_\-]+\.[a-zA-Z0-9_\-]+[a-zA-Z0-9]+$`)

func (user *User) Validate(v *revel.Validation) {
	ValidateEmail(v, user.Email).Key("user.Email")
	// ValidatePassword(v, user.Password).Key("user.Password")
	v.Check(user.Name, revel.Required{}, revel.MaxSize{100})
}

// func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
// 	return v.Check(password,
// 		revel.Required{},
// 		revel.MaxSize{15},
// 		revel.MinSize{5},
// 	)
// }

func ValidateEmail(v *revel.Validation, email string) *revel.ValidationResult {
	return v.Check(email,
		revel.Required{},
		revel.Match{emailRegexp},
	)
}
