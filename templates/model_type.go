package templates

var ModelTypeTemplate = `package models

import "github.com/jinzhu/gorm"

type Model struct {
	Db *gorm.DB
}

func NewModel(db *gorm.DB) Model{
	/* AUTO_MIGRATION_PLACEHOLDER */
	autoMigrate(db)
	return Model{Db: db}
}
`
