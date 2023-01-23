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
