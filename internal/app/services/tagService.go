package services

import (
	"Gallery/internal/app/models"
	"Gallery/internal/app/utils"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

type TagService struct {
}

var converter = utils.Converter{}

func (ts *TagService) TagsByLetters(db *gorm.DB, r *http.Request) []string {
	var (
		tagInCome  string
		tags       []models.Tags
		tagsToSend []string
	)
	err := json.NewDecoder(r.Body).Decode(&tagInCome)
	if tagInCome == " " || tagInCome == "" {
		return nil
	}
	if err != nil {
		log.Fatalln(err)
	}
	db.Select("name").Where("name ~ ?", tagInCome).Find(&tags)
	for _, v := range tags {
		tagsToSend = append(tagsToSend, v.Name)
	}
	return tagsToSend
}

func (ts *TagService) FindPostByNameTag(db *gorm.DB, r *http.Request) interface{} {
	var data []map[string]interface{}
	tagStr := chi.URLParam(r, "tagNames")
	tagNames := strings.Split(tagStr, "+")
	db.Raw("select "+
		"posts.id   as id,"+
		"posts.url  as url,"+
		"tags.id    as tag_id,"+
		"tags.name  as tag_name,"+
		"users.id   as user_id,"+
		"users.name as user_name "+
		"from posts "+
		"join users on users.id = posts.user_id "+
		"join posts_tags as pt on posts.id = pt.posts_id "+
		"join tags on tags.id = pt.tags_id "+
		"where pt.posts_id in (select p.id "+
		"from posts p "+
		"join posts_tags pt on p.id = pt.posts_id "+
		"join tags t on t.id = pt.tags_id "+
		"where t.name ~ ? )",
		tagNames).Scan(&data)
	if len(data) == 0 {
		return http.StatusNoContent
	}
	return c.TransformFromSearchByTag(data)
}
