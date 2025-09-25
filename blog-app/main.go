package main

import (
	"blog-app/Model"
	"blog-app/scopes"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "host=localhost user=postgres password=1289 dbname=Blog_App port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err != nil {
		panic("Failed to ping database")
	}
	db.AutoMigrate(&Model.User{}, &Model.Post{}) // Auto-migrate User and Post models
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&Model.User{Name: "Ermias Abebe", Email: "Ermias@gmail.com", Password: "12345"}).Error; err != nil {
			return err
		}
		return nil
	})
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&Model.Post{Title: "First Post", UserID: 1}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic("Transaction failed")
	}

	var user Model.User
	const userID = 1 // Use the ID of the created user
	db.Preload("Posts").Scopes(scopes.ByUserID(userID)).First(&user)
	fmt.Printf("User: %s has %d posts\n", user.Name, len(user.Posts))
	db.Create(&Model.Post{Title: "Second Post", Content: "Another post", UserID: userID})
	var posts []Model.Post
	db.Scopes(scopes.ByUserID(userID)).Find(&posts)
	fmt.Println("Posts:", len(posts))
	db.Delete(&posts[0]) // Delete the first post
	fmt.Println("After delete, posts:", len(posts)-1)

	fmt.Println("Mini project completed!")

}
