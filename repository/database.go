package repository

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

var (
	DBConnect *gorm.DB
)
