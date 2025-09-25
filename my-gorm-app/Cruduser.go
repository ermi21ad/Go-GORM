package main

import (
	"fmt"
	"my-gorm-app/Model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CrudUser() {
	dsn := "host=localhost user=postgres password=1289 dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err != nil {
		panic("Failed to ping database")

	}
	// Create
	db.Create(&Model.User{Name: "Ermias", Email: "ermi@example.com"})
	fmt.Println("User Created Successfully")
	// Read
	var user Model.User
	db.First(&user, "email=?", "ermi@example.com")
	fmt.Println("User Read Successfully:", user)

	// Update
	db.Model((&user)).Update("Name", "Ermias Abebe")

	fmt.Println("User Updated Successfully:", user)

}
func main() {
	CrudUser()
	fmt.Println("Operation Completed")
}
