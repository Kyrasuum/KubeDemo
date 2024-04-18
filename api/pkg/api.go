package pkg

import (
	"api/internal"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// http host
func Host() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wrapper(w, r, internal.GetData)
	})
	fmt.Printf("server started on port: 8082\n")
	err := http.ListenAndServe(":8082", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
	return nil
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
