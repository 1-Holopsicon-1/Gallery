package models

type Tags struct {
	ID    uint     `gorm:"primaryKey;autoIncrement"`
	Name  string   `gorm:"varchar(120)"`
	Posts []*Posts `gorm:"many2many:posts_tags;"`
}
