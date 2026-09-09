package main

import (
	"bytes"
	"context"
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/yuin/goldmark"
)

//go:embed templates static content.md
var content embed.FS

// renderMarkdown renders the whole document to HTML in one pass.
func renderMarkdown(src string) (string, error) {
	md := goldmark.New()
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

func newHandler() http.Handler {
	pageHTML, err := renderMarkdown(loadMarkdown())
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.ParseFS(content, "templates/*.html"))

	staticFS, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		data := struct{ FullHTML template.HTML }{template.HTML(pageHTML)}
		if err := tmpl.ExecuteTemplate(w, "shell", data); err != nil {
			log.Printf("render shell: %v", err)
		}
	})

	staticHandler := http.StripPrefix("/static", http.FileServerFS(staticFS))
	mux.Handle("GET /static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=604800")
		staticHandler.ServeHTTP(w, r)
	}))

	return mux
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
