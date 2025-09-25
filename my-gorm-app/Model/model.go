package Model

type User struct {
	ID    uint   `gorm:"primartKey`
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"type:varchar(100);unique"`
}
type Post struct {
	ID     uint
	Title  string `gorm:"type:varchar(100)"`
	UserID uint
}
