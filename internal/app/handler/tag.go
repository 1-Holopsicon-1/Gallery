package handler

import (
	"Gallery/internal/app/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) findByTag(w http.ResponseWriter, r *http.Request) {
	var tag models.Tags
	tagName := chi.URLParam(r, "tagName")
	h.DB.Preload("Posts").Where("name ~~ ?", tagName).Find(&tag)
	fmt.Println()
	//TODO: Load all tags related to post
	json.NewEncoder(w).Encode(&tag)
}
