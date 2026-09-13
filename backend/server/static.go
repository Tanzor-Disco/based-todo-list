package server

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static
var staticFS embed.FS

func handleStatic(w http.ResponseWriter, r *http.Request) {
	path := "static" + r.URL.Path
	indexPath := "static/index.html"

	if _, err := fs.Stat(staticFS, path); err == nil {
		http.ServeFileFS(w, r, staticFS, path)
		return
	}
	http.ServeFileFS(w, r, staticFS, indexPath)
}
