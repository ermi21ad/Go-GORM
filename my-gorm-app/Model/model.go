package Model

type User struct {
	ID    uint   `gorm:"primartKey`
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"type:varchar(100);unique"`
}
type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(100)"`
	Price float64
}
