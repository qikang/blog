package router

import (
	"blog/config"
	"blog/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router
func NewRouter(handler *handlers.BlogHandler, cfg *config.Config) *mux.Router {
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

	// 文章中相对路径引用的图片等资源,从 posts 目录提供(必须与 Dockerfile 中 COPY 的目录一致)
	router.PathPrefix("/posts/").Handler(http.StripPrefix("/posts/", http.FileServer(http.Dir(cfg.PostsDir))))

	return router
}