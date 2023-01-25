package utils

import (
	"Gallery/internal/app/dto"
	"Gallery/internal/app/models"
)

type Converter struct {
}

func (c *Converter) UserToDtoNP(u models.User) dto.UserDtoNP {
	var uDto dto.UserDtoNP
	uDto.ID = u.ID
	uDto.Name = u.Name
	return uDto

}

func (c *Converter) TagToDto(t models.Tags) dto.TagDto {
	var tDto dto.TagDto
	tDto.ID = t.ID
	tDto.Name = t.Name
	return tDto
}

func (c *Converter) PostToDto(p models.Posts) dto.PostDto {
	var pDto dto.PostDto
	pDto.ID = p.ID
	pDto.Url = p.Url
	for _, v := range p.Tags {
		pDto.Tags = append(pDto.Tags, c.TagToDto(*v))
	}
	pDto.User = c.UserToDtoNP(p.User)
	return pDto
}

func (c *Converter) TransformFromSearchByTag(data []map[string]interface{}) []dto.PostDto {
	var pDtos []dto.PostDto
	pDto := dto.PostDto{
		ID:  uint(data[0]["id"].(int64)),
		Url: data[0]["url"].(string),
		User: dto.UserDtoNP{
			ID:   uint(data[0]["user_id"].(int64)),
			Name: data[0]["user_name"].(string)},
	}
	for i, v := range data {
		if pDto.ID == uint(v["id"].(int64)) {
			tag := dto.TagDto{
				ID:   uint(v["tag_id"].(int64)),
				Name: v["tag_name"].(string)}
			pDto.Tags = append(pDto.Tags, tag)
		} else {
			pDtos = append(pDtos, pDto)
			pDto = dto.PostDto{
				ID:   uint(v["id"].(int64)),
				Url:  v["url"].(string),
				Tags: []dto.TagDto{},
				User: dto.UserDtoNP{
					ID:   uint(v["user_id"].(int64)),
					Name: v["user_name"].(string),
				},
			}
			tag := dto.TagDto{
				ID:   uint(v["tag_id"].(int64)),
				Name: v["tag_name"].(string)}
			pDto.Tags = append(pDto.Tags, tag)
		}
		if i == len(data)-1 {
			pDtos = append(pDtos, pDto)
		}
	}
	return pDtos
}
