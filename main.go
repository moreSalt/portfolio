package portfolio

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
)

//go:embed templates static content.md
var content embed.FS

var sectionRe = regexp.MustCompile(`(?m)^##\s+`)

type section struct {
	Slug string
	HTML string
}

type shellData struct {
	FirstFrag string
	FullHTML  template.HTML
}

func slugify(s string) string {
	var b strings.Builder
	lastDash := true // trims a leading dash
	for _, r := range strings.ToLower(strings.TrimSpace(s)) {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			b.WriteRune(r)
			lastDash = false
		} else if !lastDash {
			b.WriteByte('-')
			lastDash = true
		}
	}
	return strings.TrimRight(b.String(), "-")
}

// splitSections splits markdown into an intro chunk (everything before the
// first "## ") plus one chunk per "## " heading, in document order.
func splitSections(src string) []string {
	locs := sectionRe.FindAllStringIndex(src, -1)
	if len(locs) == 0 {
		return []string{src}
	}
	var chunks []string
	if locs[0][0] > 0 {
		chunks = append(chunks, src[:locs[0][0]])
	}
	for i, loc := range locs {
		end := len(src)
		if i+1 < len(locs) {
			end = locs[i+1][0]
		}
		chunks = append(chunks, src[loc[0]:end])
	}
	return chunks
}

func chunkSlug(chunk, fallback string) string {
	line := strings.SplitN(chunk, "\n", 2)[0]
	line = strings.TrimSpace(strings.TrimLeft(line, "#"))
	if s := slugify(line); s != "" {
		return s
	}
	return fallback
}

// buildSections renders markdown to HTML sections chained via htmx triggers.
func buildSections(src string) ([]section, error) {
	md := goldmark.New()
	chunks := splitSections(src)
	sections := make([]section, 0, len(chunks))
	for i, chunk := range chunks {
		var buf bytes.Buffer
		if err := md.Convert([]byte(chunk), &buf); err != nil {
			return nil, fmt.Errorf("section %d: %w", i, err)
		}
		html := buf.String()
		// external links open in a new tab
		html = strings.ReplaceAll(html, `<a href="http`, `<a target="_blank" rel="noopener" href="http`)
		slug := "intro"
		if i > 0 || strings.HasPrefix(strings.TrimSpace(chunk), "## ") {
			slug = chunkSlug(chunk, fmt.Sprintf("section-%d", i))
		}
		sections = append(sections, section{Slug: slug, HTML: html})
	}
	for i := range sections {
		var b strings.Builder
		fmt.Fprintf(&b, `<section id="%s">`, sections[i].Slug)
		b.WriteString(sections[i].HTML)
		if i+1 < len(sections) {
			fmt.Fprintf(&b, `<div hx-get="/frag/%s" hx-trigger="typed" hx-swap="outerHTML"></div>`, sections[i+1].Slug)
		}
		b.WriteString(`</section>`)
		sections[i].HTML = b.String()
	}
	return sections, nil
}

func loadMarkdown() string {
	// a content.md on disk overrides the embedded one
	raw, err := os.ReadFile("content.md")
	if err != nil {
		raw, err = content.ReadFile("content.md")
		if err != nil {
			log.Fatal("no content.md found on disk or embedded")
		}
	}
	return string(raw)
}

// NewHandler builds the site's HTTP handler: the shell page, the htmx
// fragment endpoints, and the static file server.
func NewHandler() http.Handler {
	sections, err := buildSections(loadMarkdown())
	if err != nil {
		log.Fatal(err)
	}

	fragments := make(map[string]string, len(sections))
	var fullHTML strings.Builder
	for _, s := range sections {
		fragments[s.Slug] = s.HTML
		fullHTML.WriteString(s.HTML)
	}

	tmpl := template.Must(template.ParseFS(content, "templates/*.html"))

	staticFS, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		data := shellData{
			FirstFrag: sections[0].Slug,
			FullHTML:  template.HTML(fullHTML.String()),
		}
		if err := tmpl.ExecuteTemplate(w, "shell", data); err != nil {
			log.Printf("render shell: %v", err)
		}
	})

	mux.HandleFunc("GET /frag/{slug}", func(w http.ResponseWriter, r *http.Request) {
		html, ok := fragments[r.PathValue("slug")]
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, html)
	})

	staticHandler := http.StripPrefix("/static", http.FileServerFS(staticFS))
	mux.Handle("GET /static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=604800")
		staticHandler.ServeHTTP(w, r)
	}))

	return mux
}
