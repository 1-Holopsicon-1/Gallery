package models

type User struct {
	ID       uint    `gorm:"primaryKey;autoIncrement"`
	Name     string  `gorm:"varchar(120)"`
	Password string  `gorm:"varchar(32)"`
	Posts    []Posts `gorm:"constraint:OnDelete:CASCADE;"`
}
