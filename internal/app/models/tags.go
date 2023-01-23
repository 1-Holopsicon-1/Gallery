package models

type Tags struct {
	ID    uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string   `gorm:"varchar(120)" json:"tag_name"`
	Posts []*Posts `gorm:"many2many:posts_tags;"`
}
