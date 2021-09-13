package repository

import (
	"github.com/jinzhu/gorm"

	"user-service/src/core"
)

func getDB() *gorm.DB {
	return core.GetDB()
}
