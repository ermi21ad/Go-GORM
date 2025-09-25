package migrations

import (
	"gorm.io/gorm"
)
func MigrateUpposts(db *gorm.DB) {
	err := db.Exec(`CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100),	
		user_id INTEGER REFERENCES users(id)
	)`).Error
	if err != nil {
		panic("Failed to create posts table: " + err.Error())
	}
}

func MigrateDownposts(db *gorm.DB) {
	err := db.Exec(`DROP TABLE IF EXISTS posts`).Error
	if err != nil {
		panic("Failed to drop posts table: " + err.Error())
	}
}