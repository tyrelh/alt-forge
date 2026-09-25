package main

import (
	"embed"
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"
)

//go:embed templates/*.html static/*
var assets embed.FS
var pages = template.Must(template.ParseFS(assets, "templates/*.html"))

type pageData struct {
	Ancestries []string
	Selected   string
	Names      []Name
}

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServer(http.FS(assets)))
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		ancestry := r.URL.Query().Get("ancestry")
		if ancestry == "" {
			ancestry = "Human"
		}
		names, ok := generate(ancestry, 6)
		if !ok {
			http.Error(w, "Unknown ancestry", http.StatusBadRequest)
			return
		}
		data := pageData{ancestries, ancestry, names}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Vary", "HX-Request")
		view := "page"
		if r.Header.Get("HX-Request") == "true" {
			view = "names"
		}
		if err := pages.ExecuteTemplate(w, view, data); err != nil {
			log.Print(err)
		}
	})
	return mux
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8088", "HTTP listen address")
	flag.Parse()
	server := &http.Server{Addr: *addr, Handler: handler(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 60 * time.Second}
	log.Printf("AltForge is ready at http://%s", *addr)
	log.Fatal(server.ListenAndServe())
}
