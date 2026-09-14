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
	r := router.NewRouter(handler)

	// Add logging middleware
	loggedRouter := logRequest(r)

	// Start server
	addr := cfg.Port
	log.Printf("Starting server on %s", addr)
	log.Printf("Site: %s", siteName)
	log.Printf("Posts directory: %s", cfg.PostsDir)

	if err := http.ListenAndServe(addr, loggedRouter); err != nil {
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
