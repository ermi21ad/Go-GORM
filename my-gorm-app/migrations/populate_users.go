package migrations

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100"`
	Email     string    `gorm:"size:100;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // Automatically set to current time
}

func PopulateUsers(db *gorm.DB) {
	users := []User{
		{Name: "Abebe Kedir", Email: "abebe.kedir@example.com"},
		{Name: "Fatuma Mohammed", Email: "fatuma.mohammed@example.com"},
		{Name: "Hana Tesfaye", Email: "hana.tesfaye@example.com"},
		{Name: "Kebede Assefa", Email: "kebede.assefa@example.com"},
		{Name: "Marta Beshah", Email: "marta.beshah@example.com"},
		{Name: "Nuru Ayalew", Email: "nuru.ayalew@example.com"},
		{Name: "Selamawit Berhe", Email: "selamawit.berhe@example.com"},
		{Name: "Tadesse Dibaba", Email: "tadesse.dibaba@example.com"},
		{Name: "Yemane Dagnachew", Email: "yemane.dagnachew@example.com"},
		{Name: "Zahra Abdi", Email: "zahra.abdi@example.com"},
	}

	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			panic("Failed to populate users table: " + err.Error())
		}
	}
}
