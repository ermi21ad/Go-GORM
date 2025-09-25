package migrations

import "gorm.io/gorm"

func MigrateUpposts(db *gorm.DB) {
	err := db.Exec(`CREAT TABLE IF NOT EXISTS posts(
	ID SERIAL PRIMARY KEY,
	Title VARCHAR(100),
	Content VARCHAR(100) ,
	UserID VARCHAR(100),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)`).Error
	if err != nil {
		panic("Failed to create posts table: " + err.Error())
	}
}
func MigrateDownposts(db *gorm.DB) {
	err := db.Exec(`DROP TABLE ID EXISTS posts`).Error
	if err != nil {
		panic("Failed to drop posts table: " + err.Error())
	}
}
