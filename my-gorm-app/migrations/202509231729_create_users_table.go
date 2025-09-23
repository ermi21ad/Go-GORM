package migrations

import (
	"gorm.io/gorm"
)

func MigrateUp(db *gorm.DB) {
	err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100) UNIQUE,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    )`).Error
	if err != nil {
		panic("Failed to create users table: " + err.Error())
	}
}

func MigrateDown(db *gorm.DB) {
	err := db.Exec(`DROP TABLE IF EXISTS users`).Error
	if err != nil {
		panic("Failed to drop users table: " + err.Error())
	}
}
