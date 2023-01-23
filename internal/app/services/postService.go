package services

import (
	"Gallery/internal/app/dto"
	"Gallery/internal/app/models"
	"Gallery/internal/app/utils"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

var c utils.Converter

type PostService struct {
}

func (ps *PostService) FindById(db *gorm.DB, id string) dto.PostDto {
	var (
		post    models.Posts
		postDto dto.PostDto
	)
	db.Preload("Tags").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name")
	}).Find(&post, id)
	postDto = c.PostToDto(post)
	fmt.Println(postDto)

	return postDto
}

func (ps *PostService) GetAll(db *gorm.DB) []dto.PostDto {
	var (
		posts    []models.Posts
		postDtos []dto.PostDto
	)
	db.Preload("Tags").Preload("User",
		func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).Find(&posts)
	for _, v := range posts {
		postDtos = append(postDtos, c.PostToDto(v))
	}
	return postDtos
}

func (ps *PostService) CreatePost(db *gorm.DB, r *http.Request) int {
	var post models.Posts
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatalln(err)
		return http.StatusBadRequest
	}

	if post.Url == "" {
		return http.StatusNoContent
	}
	for index, value := range post.Tags {
		var tag models.Tags
		db.FirstOrCreate(&tag, models.Tags{Name: value.Name})
		post.Tags[index].ID = tag.ID
	}

	db.FirstOrCreate(&post.User)
	db.Create(&post)

	return http.StatusCreated
}

func (ps *PostService) Delete(db *gorm.DB, id string) int {
	err := db.Delete(&models.Posts{}, id).Error
	if err != nil {
		log.Fatalln(err)
		return http.StatusBadRequest
	}
	return http.StatusAccepted
}

func (ps *PostService) EditTags(db *gorm.DB, r *http.Request) int {
	var tags []models.Tags
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&tags)
	if err != nil {
		log.Fatalln(err)
		return http.StatusBadRequest
	}

	for i, v := range tags {
		db.FirstOrCreate(&tags[i], models.Tags{Name: v.Name})
	}
	uid, _ := strconv.ParseUint(id, 10, 0)
	err = db.Model(&models.Posts{ID: uint(uid)}).Association("Tags").Replace(&tags)
	if err != nil {
		log.Fatalln(err)
		return http.StatusTeapot
	}
	return http.StatusAccepted
}
