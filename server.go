package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/srohatgi/wtwt/weather"
	"github.com/vharitonsky/iniflags"
	"io"
	"log"
	"net/http"
)

var (
	port = flag.Int("port", 8000, "port# to listen")
)

func makeHandler(fn func(*http.Request, map[string]string) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("calling for %s, vars: %s", r.URL.RequestURI(), vars)
		if x, err := fn(r, vars); err == nil {
			j, _ := json.Marshal(x)
			io.WriteString(w, string(j))
		} else {
			io.WriteString(w, fmt.Sprintf("error: %v", err))
		}
	}
}

func main() {
	iniflags.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/conditions", makeHandler(func(r *http.Request, vars map[string]string) (interface{}, error) {
		m := r.URL.Query()
		if _, ok := m["ip"]; ok {
			return weather.GetConditions(100)
		}
		return struct{}{}, errors.New("unimplemented")
	}))
	log.Printf("listening on %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), r))
}
