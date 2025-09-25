package Model

import "gorm.io/gorm"

type User struct {
	Id       uint   `gorm:"PrimartKey"`
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type :varchar(100);unique"`
	Password string `gorm:"type:varchar(100)"`
	Posts    []Post
}

type Post struct {
	Id      uint
	Title   string `gorm:"type:varchar(100)"`
	Content string `gorm:"type:text"`
	UserID  uint
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Password = "hashed_" + u.Password
	return nil
}
