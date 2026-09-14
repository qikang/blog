package router

import (
	"blog/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router
func NewRouter(handler *handlers.BlogHandler) *mux.Router {
	router := mux.NewRouter()

	// API routes - more specific routes first
	router.HandleFunc("/", handler.IndexHandler)
	router.HandleFunc("/page/{page}", handler.IndexHandler)
	router.HandleFunc("/post/{slug:.*}", handler.PostHandler)
	router.HandleFunc("/archives", handler.ArchivesHandler)
	router.HandleFunc("/tags", handler.TagsHandler)
	router.HandleFunc("/tag/{tag}", handler.TagHandler)
	router.HandleFunc("/about", handler.AboutHandler)

	// Static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	return router
}
