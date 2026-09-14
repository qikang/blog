package handlers

import (
	"blog/models"
	"net/http"

	"github.com/gorilla/mux"
)

// TagsHandler handles the tags list page
func (h *BlogHandler) TagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := h.mdService.GetAllTags()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		SiteName string
		Tags     []string
	}{
		SiteName: h.siteName,
		Tags:     tags,
	}

	h.renderTemplate(w, r, "tags.html", data)
}

// TagHandler handles posts by tag page
func (h *BlogHandler) TagHandler(w http.ResponseWriter, r *http.Request) {
	tag := mux.Vars(r)["tag"]
	if tag == "" {
		http.NotFound(w, r)
		return
	}

	posts, err := h.mdService.GetPostsByTag(tag)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		SiteName string
		Tag      string
		Posts    []models.Post
	}{
		SiteName: h.siteName,
		Tag:      tag,
		Posts:    posts,
	}

	h.renderTemplate(w, r, "tag.html", data)
}
