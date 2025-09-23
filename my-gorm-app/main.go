package main

import (
	"fmt"
	"my-gorm-app/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
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
	migrations.MigrateDown(db)
	migrations.MigrateUp(db)

	fmt.Println("Database connected and migrations applied successfully")
}
