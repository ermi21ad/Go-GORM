package main

import (
	"fmt"
	"my-gorm-app/migrations" // Ensure this path is correct

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=1289 dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err != nil {
		panic("Failed to ping database: " + err.Error())
	}

	// migrations.PopulateUsers(db)
	migrations.PopulatePosts(db)

	fmt.Println("Database connected, migrations applied, and populated successfully")
}
