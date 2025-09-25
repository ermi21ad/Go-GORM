package migrations

import (
	"gorm.io/gorm"
)

type Post struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"type:varchar(100)"`
	UserID uint   `gorm:"index"` // Foreign key for User
}

func PopulatePosts(db *gorm.DB) {
	posts := []Post{
		{Title: "First Post", UserID: 5},
		{Title: "Second Post", UserID: 6},
		{Title: "Third Post", UserID: 10},
		{Title: "Fourth Post", UserID: 8},
		{Title: "Fifth Post", UserID: 9},
	}

	for _, post := range posts {
		err := db.Create(&post).Error
		if err != nil {
			panic("Failed to populate posts table: " + err.Error())
		}
	}
}
