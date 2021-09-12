package database

import (
	"github.com/FerRiosCosta/go-rest-api-crud/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates out database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
