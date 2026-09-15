package router

import (
	"blog/assets"
	"blog/config"
	"blog/handlers"
	"io/fs"
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

	// 静态资源(templates/css/js/images)从 embed.FS 提供,
	//先 fs.Sub 裁掉 "static" 前缀,让 FileServer 看到的就是 static/ 下的相对路径。
	staticFS, err := fs.Sub(assets.StaticFS, "static")
	if err != nil {
		panic("assets.StaticFS sub static failed: " + err.Error())
	}
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))),
	)

	// 文章中相对路径引用的图片等资源,从 posts 目录提供(markdown 文件目录,运行期可变)。
	router.PathPrefix("/posts/").Handler(
		http.StripPrefix("/posts/", http.FileServer(http.Dir(cfg.PostsDir))),
	)

	return router
}
