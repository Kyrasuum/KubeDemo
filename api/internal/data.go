package internal

import (
	"embed"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
)

//go:embed data
var data embed.FS

// http data request handler
func GetData(w http.ResponseWriter, r *http.Request, args ...interface{}) {
	//create request path
	path := r.URL.Path
	path = "data" + path

	//get page content
	p, err := data.ReadFile(path)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//get content type
	ctype := mime.TypeByExtension(filepath.Ext(path))
	if ctype == "" {
		ctype = http.DetectContentType(p)
	}

	//write content
	w.Header().Add("Content-Type", ctype)
	w.WriteHeader(http.StatusOK)
	w.Write(p)
}
