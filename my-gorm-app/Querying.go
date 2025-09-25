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
	var users []Model.User
	db.Where("name LIKE ?", "%A%").Select("name, email").Order("name DESC").Limit(2).Offset(1).Find(&users)
	for _, user := range users {
		fmt.Printf("Combined query user: %s (Email: %s)\n", user.Name, user.Email)
	}
}
