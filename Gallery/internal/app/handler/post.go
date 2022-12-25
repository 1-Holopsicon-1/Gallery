package handler

import (
	"Gallery/internal/app/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func (h *Handler) findPostById(w http.ResponseWriter, r *http.Request) {
	type resultSt struct {
		ID       uint
		Url      string
		Tags     []*models.Tags
		UserName string
	}
	id := chi.URLParam(r, "id")
	var post models.Posts
	var user map[string]interface{}
	fmt.Println(h.DB.Find(&post, id).Joins("join users on user_id=users.id").Scan(&user).Preload("Tags"))
	result := resultSt{ID: post.ID, Url: post.Url, Tags: post.Tags, UserName: user["name"].(string)}
	fmt.Println(user)
	json.NewEncoder(w).Encode(&result)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	var posts []models.Posts
	h.DB.Find(&posts)
	h.DB.Preload("Tags").Find(&posts)
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		log.Fatalln(err)
	}
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	var post models.Posts
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		//
		panic(err)
	}
	h.DB.Create(&post)
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.DB.Delete(&models.Posts{}, id)
}

func (h *Handler) editTags(w http.ResponseWriter, r *http.Request) {

}
