package internal

import (
	"app/pkg"
	"embed"
	"errors"
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:embed pages/dist
var pages embed.FS

// http host
func Host() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wrapper(w, r, getPage)
	})
	fmt.Printf("server started on port: 8081\n")
	err := http.ListenAndServe(":8081", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
	return nil
}

// http page request handler
func getPage(w http.ResponseWriter, r *http.Request, args ...interface{}) {
	//create request path
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}

	//get page content
	p := []byte{}
	var err error
	if !strings.HasPrefix(path, "/api/") {
		path = "pages/dist" + path
		p, err = pages.ReadFile(path)
		if err != nil {
			fmt.Printf("%+v\n", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
	} else {
		p, err = pkg.GetData(path[4:])
		if err != nil {
			fmt.Printf("%+v\n", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
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

// wrapper for http requests
func wrapper(
	wri http.ResponseWriter,
	req *http.Request,
	f func(http.ResponseWriter,
		*http.Request,
		...interface{},
	),
	args ...interface{}) {
	//seperation for clarity
	addCorsHeader(wri)
	if req.Method == "OPTIONS" {
		wri.WriteHeader(http.StatusOK)
		return
	} else {
		f(wri, req, args...)
	}
}

// cors hanlder for http requests
func addCorsHeader(wri http.ResponseWriter) {
	headers := wri.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
}
