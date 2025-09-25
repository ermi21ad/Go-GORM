package migrations

import "gorm.io/gorm"

func MigrateUpusers(db *gorm.DB) {
	err := db.Exec(`CREAT TABLE IF NOT EXISTS users(
	ID SERIAL PRIMARY KEY,
	Name VARCHAR(100),
	Email VARCHAR(100) UNIQUE,
	Password VARCHAR(100),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)`).Error
	if err != nil {
		panic("Failed to create users table: " + err.Error())
	}
}
func MigrateDownusers(db *gorm.DB) {
	err := db.Exec(`DROP TABLE ID EXISTS users`).Error
	if err != nil {
		panic("Failed to drop users table: " + err.Error())
	}
}
