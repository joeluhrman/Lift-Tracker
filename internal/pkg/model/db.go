package model

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

// Generic table model.
type Model struct {
	ID uint `gorm:"primaryKey;not null" json:"id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Creates/auto-migrates a local SQLite db at the specified path parameter.
func InitDB(path string) error {
	var err error
	DBConn, err = gorm.Open(sqlite.Open(path))
	if err != nil {
		return err
	}

	err = DBConn.AutoMigrate(
		&User{},
		&UserPassword{},
	)

	return err
}
