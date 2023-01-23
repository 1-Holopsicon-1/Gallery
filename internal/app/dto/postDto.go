package dto

type PostDto struct {
	ID   uint      `json:"id"`
	Url  string    `gorm:"text" json:"url"`
	Tags []TagDto  `json:"tags"`
	User UserDtoNP `json:"user"`
}
