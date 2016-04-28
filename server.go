package main

import (
	"fmt"
	"github.com/srohatgi/wtwt/weather"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	lon, lat := weather.GetLongLat(100)
	coords := fmt.Sprintf("%f, %f", lon, lat)
	io.WriteString(w, coords)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
