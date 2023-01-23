package handler

import (
	"Gallery/internal/app/services"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var postService services.PostService

func (h *Handler) findPostById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := json.NewEncoder(w).Encode(postService.FindById(h.DB, id))
	if err != nil {
		log.Fatalln(err)
	}
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(postService.GetAll(h.DB))
	if err != nil {
		log.Fatalln(err)
	}
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(postService.CreatePost(h.DB, r))
	if err != nil {
		log.Fatalln(err)
	}
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := json.NewEncoder(w).Encode(postService.Delete(h.DB, id))
	if err != nil {
		log.Fatalln(err)
	}
}

func (h *Handler) editTags(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(postService.EditTags(h.DB, r))
	if err != nil {
		log.Fatalln(err)
	}
}
