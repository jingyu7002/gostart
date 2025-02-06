package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}
