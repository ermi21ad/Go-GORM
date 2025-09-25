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
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&Model.User{Name: "Habtu", Email: "Habtu@example.com"}).Error; err != nil {
			return err
		}
		return nil
	})
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&Model.Post{Title: "eyob", UserID: 8}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic("Transaction failed")
	}
	fmt.Println("Transaction completed!")
}
