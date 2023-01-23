package models

type Posts struct {
	ID     uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Url    string  `gorm:"text" json:"url"`
	Tags   []*Tags `gorm:"many2many:posts_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"tags"`
	UserId int     `json:"user_id"`
	User   User
}
