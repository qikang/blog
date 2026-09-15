package handlers

import (
	"blog/assets"
	"blog/config"
	"blog/services"
	"html/template"
	"net/http"
	"reflect"
	"strings"
)

// BlogHandler handles blog HTTP requests
type BlogHandler struct {
	config    *config.Config
	mdService *services.MarkdownService
	siteName  string
	pageSize  int
	tmpl      *template.Template // 启动时一次性从 embed.FS 解析所有 *.html
}

// NewBlogHandler creates a new blog handler
func NewBlogHandler(cfg *config.Config, siteName string, pageSize int) *BlogHandler {
	// 一次性解析嵌入的 HTML 模板,后续请求直接复用 *Template,避免每次 ParseGlob
	tmpl := template.Must(template.ParseFS(assets.TemplatesFS, "templates/*.html"))

	return &BlogHandler{
		config:    cfg,
		mdService: services.NewMarkdownService(cfg.PostsDir),
		siteName:  siteName,
		pageSize:  pageSize,
		tmpl:      tmpl,
	}
}

// renderTemplate renders a template with the given data
func (h *BlogHandler) renderTemplate(w http.ResponseWriter, r *http.Request, tmplName string, data interface{}) {
	// Get current path
	currentPath := r.URL.Path

	// Determine current page type
	currentPage := ""
	switch {
	case currentPath == "/":
		currentPage = "home"
	case currentPath == "/archives":
		currentPage = "archives"
	case currentPath == "/tags":
		currentPage = "tags"
	case currentPath == "/about":
		currentPage = "about"
	case strings.HasPrefix(currentPath, "/tag/"):
		currentPage = "tag"
	case strings.HasPrefix(currentPath, "/post/"):
		currentPage = "post"
	}

	// Determine which page the user came from (for post pages)
	referer := r.Referer()
	fromPage := "archives" // default
	if referer != "" {
		if strings.HasPrefix(referer, r.Host) {
			path := strings.TrimPrefix(referer, "http://"+r.Host)
			path = strings.TrimPrefix(path, "https://"+r.Host)
			if path == "/" || strings.HasPrefix(path, "/page/") {
				fromPage = "home"
			} else if path == "/archives" {
				fromPage = "archives"
			}
		}
	}

	// Convert data to map and add CurrentPath
	pageData := map[string]interface{}{
		"CurrentPath": currentPath,
		"CurrentPage": currentPage,
		"FromPage":    fromPage,
	}

	if data != nil {
		// If data is already a map, merge it
		if dataMap, ok := data.(map[string]interface{}); ok {
			for k, v := range dataMap {
				pageData[k] = v
			}
		} else {
			// Use reflection to get struct fields
			v := reflect.ValueOf(data)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				t := v.Type()
				for i := 0; i < t.NumField(); i++ {
					field := t.Field(i)
					pageData[field.Name] = v.Field(i).Interface()
				}
			}
		}
	}

	if err := h.tmpl.ExecuteTemplate(w, tmplName, pageData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
