package dto

type UserDtoNP struct {
	ID   uint   `json:"id"`
	Name string `json:"user_name"`
}

type UserDtoP struct {
	UserDtoNP
	Posts []PostDto `json:"posts"`
}
