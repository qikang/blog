package handlers

import (
	"net/http"
)

// StaticHandler serves static files
func (h *BlogHandler) StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir(h.config.StaticDir)).ServeHTTP(w, r)
}
