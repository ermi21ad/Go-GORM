package Model

type User struct {
	ID    uint   `gorm:"primartKey`
	Name  string `gorm:"type:varchar(100)"`
	Age   int    `gorm:"type:int"`
	Email string `gorm:"type:varchar(100);unique"`
}
