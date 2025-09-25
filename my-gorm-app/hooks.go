package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func main() {
	dsn := "host=localhost user=postgres password=1289 dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.Create(&User{Name: "Bini", Email: "Bini@example.com", Password: "Pass@123"})
	fmt.Println("New user with hashed password created!")

}

type User struct {
	ID       uint
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100);unique"`
	Password string `gorm:"type:varchar(100)"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	rand.Seed(time.Now().UnixNano())                         // Seed the random number generator
	randomNumber := rand.Intn(9) + 1                         // Generate a random number between 1 and 9
	u.Password = u.Password + "_" + string(randomNumber+'0') // Append random number to the password
	return nil
}
