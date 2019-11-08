package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User is user model
type User struct {
	gorm.Model
	FirstName  string `gorm:"type:varchar(512)"`
	LastName   string `gorm:"type:varchar(512)"`
	Email      string `gorm:"type:varchar(512)"`
	Phone      string `gorm:"type:varchar(512)"`
	Picture    string `gorm:"type:varchar(512)"`
	IsVerified bool   `gorm:"type:tinyint"`
	MetaData   string `gorm:"type:text"`
}
