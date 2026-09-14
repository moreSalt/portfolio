package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"embed"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

//go:embed templates static content.md
var content embed.FS

// renderMarkdown renders the whole document to HTML in one pass.
func renderMarkdown(src string) (string, error) {
	md := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("github"),
			),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(src), &buf); err != nil {
		return "", err
	}
	html := buf.String()
	// external links open in a new tab
	html = strings.ReplaceAll(html, `<a href="http`, `<a target="_blank" rel="noopener" href="http`)
	return html, nil
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

func loadCSS() string {
	// a static/style.css on disk overrides the embedded one
	raw, err := os.ReadFile("static/style.css")
	if err != nil {
		raw, err = content.ReadFile("static/style.css")
		if err != nil {
			log.Fatal("no style.css found on disk or embedded")
		}
	}
	return string(raw)
}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}

func gzipHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") ||
			strings.HasSuffix(r.URL.Path, ".jsdos") ||
			strings.HasSuffix(r.URL.Path, ".zip") {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Add("Vary", "Accept-Encoding")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		next.ServeHTTP(&gzipResponseWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

func newHandler() http.Handler {
	pageHTML, err := renderMarkdown(loadMarkdown())
	if err != nil {
		log.Fatal(err)
	}
	inlineCSS := loadCSS()

	tmpl := template.Must(template.ParseFS(content, "templates/*.html"))

	staticFS, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	data := struct {
		FullHTML  template.HTML
		InlineCSS template.CSS
	}{template.HTML(pageHTML), template.CSS(inlineCSS)}

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "shell", data); err != nil {
			log.Printf("render shell: %v", err)
		}
	})

	mux.HandleFunc("GET /ski-or-die", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "game", data); err != nil {
			log.Printf("render game: %v", err)
		}
	})

	staticHandler := http.StripPrefix("/static", http.FileServerFS(staticFS))
	mux.Handle("GET /static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		staticHandler.ServeHTTP(w, r)
	}))

	return gzipHandler(mux)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           newHandler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		log.Printf("listening on http://localhost:%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown: %v", err)
	}
}
