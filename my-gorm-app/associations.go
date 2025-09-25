package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"my-gorm-app/Model"
)

func main() {
	dsn := "host=localhost user=postgres password=1289 dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	var user User
	db.Preload("Posts").First(&user, 1)
	fmt.Println("User:", user.Name, "has", len(user.Posts), "posts")
}

type User struct {
	ID    uint
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"type:varchar(100);unique"`
	Posts []Post
}

type Post struct {
	ID     uint
	Title  string `gorm:"type:varchar(100)"`
	UserID uint
}
