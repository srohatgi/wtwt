package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/srohatgi/wtwt/weather"
	"io"
	"log"
	"net/http"
)

func makeHandler(fn func(*http.Request, map[string]string) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("calling for %s, vars: %s", r.URL.RequestURI(), vars)
		if x, err := fn(r, vars); err != nil {
			io.WriteString(w, fmt.Sprintf("error: %v", err))
		} else {
			io.WriteString(w, fmt.Sprintf("%v", x))
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/conditions", makeHandler(func(r *http.Request, vars map[string]string) (interface{}, error) {
		return weather.GetConditions(uint32(100)), nil
	}))
	log.Printf("listening on 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
