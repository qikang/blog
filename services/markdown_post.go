package services

import (
	"blog/models"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// imgSrcRe 匹配 <img> 标签中的 src 属性,用于给 markdown 中相对路径的图片加上 /posts/<slug>/ 前缀
var imgSrcRe = regexp.MustCompile(`(<img\s+[^>]*?\bsrc=")([^"]+)"`)

// rewriteImgSrcs 把渲染后的 HTML 中所有 <img src="..."> 的相对路径加上 prefix。
// 绝对 URL(http/https/根路径/data:)保持不变,避免重复加前缀或破坏外链。
func rewriteImgSrcs(htmlStr, prefix string) string {
	return imgSrcRe.ReplaceAllStringFunc(htmlStr, func(match string) string {
		sub := imgSrcRe.FindStringSubmatch(match)
		if len(sub) < 3 {
			return match
		}
		src := sub[2]
		if src == "" || strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") ||
			strings.HasPrefix(src, "/") || strings.HasPrefix(src, "data:") {
			return match
		}
		return sub[1] + prefix + src + `"`
	})
}

// GetPostBySlug returns a single post by its slug
func (s *MarkdownService) GetPostBySlug(slug string) (*models.Post, error) {
	// Search for the file in all subdirectories
	var foundFile string
	err := filepath.Walk(s.postsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}
		// Get relative path from posts directory
		relPath, err := filepath.Rel(s.postsDir, path)
		if err != nil {
			return nil
		}
		// Remove .md extension and convert to forward slashes for comparison
		relPath = strings.TrimSuffix(relPath, ".md")
		relPath = filepath.ToSlash(relPath)
		// Check if the relative path matches the slug
		if relPath == slug {
			foundFile = path
			return filepath.SkipAll
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if foundFile == "" {
		return nil, os.ErrNotExist
	}
	return s.ParsePost(foundFile)
}

// ParsePost parses a markdown file into a Post struct
func (s *MarkdownService) ParsePost(filename string) (*models.Post, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Get relative path from posts directory
	relPath, err := filepath.Rel(s.postsDir, filename)
	if err != nil {
		relPath = filename
	}

	// Extract category (subdirectory) and slug
	dir := filepath.Dir(relPath)
	basename := strings.TrimSuffix(filepath.Base(filename), ".md")

	var slug, category string
	if dir == "." {
		slug = basename
		category = ""
	} else {
		slug = filepath.ToSlash(filepath.Join(dir, basename))
		category = strings.TrimPrefix(dir, "./")
	}

	post := &models.Post{
		Slug:     slug,
		Category: category,
	}

	// Parse frontmatter
	lines := strings.Split(string(content), "\n")
	inFrontmatter := false
	var markdownContent []string

	for i, line := range lines {
		if i == 0 && strings.TrimSpace(line) == "---" {
			inFrontmatter = true
			continue
		}
		if inFrontmatter && strings.TrimSpace(line) == "---" {
			inFrontmatter = false
			continue
		}

		if inFrontmatter {
			// Parse frontmatter key-value pairs
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				switch key {
				case "title":
					post.Title = value
				case "date":
					// 使用本地时区解析,避免跨时区部署时日期偏移一天
					if t, err := time.ParseInLocation("2006-01-02", value, time.Local); err == nil {
						post.Date = t
					}
				case "author":
					post.Author = value
				case "tags":
					// Remove surrounding brackets if present
					value = strings.Trim(value, "[]")
					post.Tags = strings.Split(value, ",")
					for i := range post.Tags {
						post.Tags[i] = strings.TrimSpace(post.Tags[i])
					}
				case "summary":
					post.Summary = value
				}
			}
		} else {
			markdownContent = append(markdownContent, line)
		}
	}

	// If no title from frontmatter, use filename
	if post.Title == "" {
		post.Title = strings.TrimSuffix(filepath.Base(filename), ".md")
	}

	// 日期优先级:frontmatter date > 文件修改时间 > 当前时间
	// (此前实现无条件使用文件 mtime,导致 frontmatter 中的 date 字段被覆盖)
	if post.Date.IsZero() {
		if stat, err := os.Stat(filename); err == nil {
			post.Date = stat.ModTime()
		} else {
			post.Date = time.Now()
		}
	}

	// Default author
	if post.Author == "" {
		post.Author = "博客作者"
	}

	// Convert markdown to HTML
	mdContent := strings.Join(markdownContent, "\n")
	post.Content = mdContent
	// 给文章内相对路径的图片加上 /posts/<目录>/ 前缀,使之命中 router 注册的 /posts/ 静态服务。
	// 约定:图片与对应 markdown 同目录(顶层 post 对应 posts/,子目录下 post 对应 posts/<子目录>/)。
	imgPrefix := "/posts/"
	if dir != "." {
		imgPrefix = "/posts/" + strings.TrimPrefix(filepath.ToSlash(dir), "./") + "/"
	}
	post.HTML = template.HTML(rewriteImgSrcs(s.MarkdownToHTML(mdContent), imgPrefix))

	return post, nil
}
