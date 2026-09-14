package services

import (
	"bytes"
	"html"
	"html/template"
	"strings"
	"unicode"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// MarkdownService handles markdown post parsing
type MarkdownService struct {
	postsDir string
}

// NewMarkdownService creates a new markdown service
func NewMarkdownService(postsDir string) *MarkdownService {
	return &MarkdownService{
		postsDir: postsDir,
	}
}

// GetPostsDir returns the posts directory path
func (s *MarkdownService) GetPostsDir() string {
	return s.postsDir
}

// slugify converts a string to a URL-safe slug, preserving Chinese characters
func slugify(text string) string {
	// First escape HTML entities
	text = html.UnescapeString(text)
	text = strings.ToLower(text)

	var result strings.Builder
	for _, r := range text {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			result.WriteRune(r)
		} else if unicode.Is(unicode.Han, r) {
			// Keep Chinese characters
			result.WriteRune(r)
		} else if r == ' ' || r == '_' || r == '.' || r == '(' || r == ')' {
			result.WriteRune('-')
		} else if r == '-' {
			result.WriteRune('-')
		}
	}

	// Clean up multiple hyphens
	id := strings.ReplaceAll(result.String(), "--", "-")
	id = strings.Trim(id, "-")
	return id
}

// addHeadingIDs adds id attributes to headings for anchor links
func addHeadingIDs(html string) string {
	// Track heading IDs to ensure uniqueness
	headingIDs := make(map[string]int)

	// Process each heading level
	for level := 1; level <= 6; level++ {
		openTag := "<h" + string(rune('0'+level)) + ">"
		closeTag := "</h" + string(rune('0'+level)) + ">"

		for {
			start := strings.Index(html, openTag)
			if start == -1 {
				break
			}

			contentStart := start + len(openTag)
			contentEnd := strings.Index(html[contentStart:], closeTag)
			if contentEnd == -1 {
				break
			}
			contentEnd += contentStart

			text := html[contentStart:contentEnd]

			// Generate ID from heading text
			id := slugify(text)

			// Skip if ID is empty
			if id == "" {
				html = html[:start] + openTag + text + closeTag + html[contentEnd+len(closeTag):]
				continue
			}

			// Ensure uniqueness
			if count, exists := headingIDs[id]; exists {
				headingIDs[id] = count + 1
				id = id + "-" + string(rune('a'+count))
			} else {
				headingIDs[id] = 0
			}

			// Replace the heading with id
			newHeading := `<h` + string(rune('0'+level)) + ` id="` + id + `">` + text + closeTag
			html = html[:start] + newHeading + html[contentEnd+len(closeTag):]
		}
	}

	return html
}

// MarkdownToHTML converts markdown to HTML
func (s *MarkdownService) MarkdownToHTML(markdownContent string) string {
	extensions := parser.CommonExtensions | parser.Attributes
	p := parser.NewWithExtensions(extensions)

	htmlBytes := markdown.ToHTML([]byte(markdownContent), p, nil)

	// Add IDs to headings for anchor links
	htmlStr := addHeadingIDs(string(htmlBytes))

	// Wrap in a container for styling
	var buf bytes.Buffer
	buf.WriteString("<div class=\"post-content\">\n")
	buf.WriteString(htmlStr)
	buf.WriteString("</div>\n")

	return buf.String()
}

// ConvertMarkdownToHTML converts markdown to template.HTML
func ConvertMarkdownToHTML(markdownContent string) template.HTML {
	ms := &MarkdownService{}
	return template.HTML(ms.MarkdownToHTML(markdownContent))
}
