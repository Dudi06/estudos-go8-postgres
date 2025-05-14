package models

import "gorm.io/gorm"

type Books struct {
	ID      uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Autor   *string `json:"autor"`
	Título  *string `json:"titulo"`
	Editora *string `json:"editora"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
