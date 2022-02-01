package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// SetDB ...
func SetDB(d *gorm.DB) {
	db = d
}
