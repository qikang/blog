package main

import (
	"blog/config"
	"blog/handlers"
	"blog/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Get site name from environment or use default
	siteName := os.Getenv("SITE_NAME")
	if siteName == "" {
		siteName = "西康的博客"
	}

	// Get page size from environment or use default
	pageSize := 5
	if ps := os.Getenv("PAGE_SIZE"); ps != "" {
		if parsed, err := fmt.Sscanf(ps, "%d", &pageSize); parsed != 1 || err != nil {
			pageSize = 5
		}
	}

	// Create handler
	handler := handlers.NewBlogHandler(cfg, siteName, pageSize)

	// Create router
	r := router.NewRouter(handler, cfg)

	// 中间件链:路由 → 404 跳首页 → 访问日志
	finalRouter := redirect404ToRoot(logRequest(r))

	// Start server
	addr := cfg.Port
	log.Printf("Starting server on %s", addr)
	log.Printf("Site: %s", siteName)
	log.Printf("Posts directory: %s", cfg.PostsDir)

	if err := http.ListenAndServe(addr, finalRouter); err != nil {
		log.Fatal(err)
	}
}

// logRequest logs incoming requests
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// statusRecorder 包装 http.ResponseWriter,记录响应状态码(用于 404 跳转)
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// redirect404ToRoot 将所有 GET 404 请求 302 跳转到首页。
// 包括静态文件 / 文章页 / 标签页等任何路径的 404(只对首页自身豁免,避免死循环)。
func redirect404ToRoot(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		if rec.status == http.StatusNotFound &&
			r.Method == http.MethodGet &&
			r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	})
}