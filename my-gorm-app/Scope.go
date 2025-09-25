package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"my-gorm-app/Model"
)

func main() {
	Scope()
}
func Scope() {
	dsn := "host=localhost user=postgres password=1289 dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	var users []Model.User
	db.Scopes(withName("E%")).Find(&users)
	for _, user := range users {
		fmt.Printf("Scoped user: %s\n", user.Name)
	}
	var usersEmail []Model.User
	db.Scopes(withEmail("%example.com")).Find(&usersEmail)
	for _, user := range usersEmail {
		fmt.Printf("Scoped user: %s (Email: %s)\n", user.Email)
	}
}
func withEmail(newpattern string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("email LIKE ?", newpattern)
	}
}

func withName(pattern string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name LIKE ?", pattern)
	}
}
