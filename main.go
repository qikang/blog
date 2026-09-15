package main

import (
	"blog/config"
	"blog/handlers"
	"blog/router"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// loadInt 从环境变量解析整数,失败时返回 fallback。
func loadInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		var n int
		if _, err := fmt.Sscanf(v, "%d", &n); err == nil {
			return n
		}
		log.Printf("环境变量 %s=%q 不是合法整数,使用默认值 %d", key, v, fallback)
	}
	return fallback
}

func main() {
	// 1. 先加载环境变量作为"配置基线"
	cfg := config.Load()
	siteName := os.Getenv("SITE_NAME")
	if siteName == "" {
		siteName = "西康的博客"
	}
	pageSize := loadInt("PAGE_SIZE", 5)

	// 2. 命令行 flag 覆盖环境变量(优先级: flag > env > default)
	flag.StringVar(&cfg.PostsDir, "posts", cfg.PostsDir, "Markdown 文章目录(支持相对/绝对路径)")
	flag.StringVar(&cfg.Port, "port", cfg.Port, "HTTP 监听地址,例如 :8083")
	flag.StringVar(&siteName, "site-name", siteName, "站点名称")
	flag.IntVar(&pageSize, "page-size", pageSize, "首页每页文章数")
	flag.Parse()

	// 3. 构造 handler / router
	handler := handlers.NewBlogHandler(cfg, siteName, pageSize)
	r := router.NewRouter(handler, cfg)

	// 4. 中间件链: 路由 → 404 跳首页 → 访问日志
	finalRouter := redirect404ToRoot(logRequest(r))

	// 5. 启动服务
	log.Printf("Starting server on %s", cfg.Port)
	log.Printf("Site: %s", siteName)
	log.Printf("Posts directory: %s", cfg.PostsDir)

	if err := http.ListenAndServe(cfg.Port, finalRouter); err != nil {
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
