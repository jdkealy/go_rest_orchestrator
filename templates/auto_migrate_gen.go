package templates


var AutoMigrateTemplate = `package models

import "github.com/jinzhu/gorm"

func autoMigrate(db *gorm.DB) {
	/* AUTO_MIGRATE_GEN */
}`