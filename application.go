package main

import (
	"net/http"
	"image/gif"
	"fmt"
)

func fetchGif(path string) *gif.GIF {
	resp, err := http.Get(path)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	g, err := gif.DecodeAll(resp.Body)
	if err != nil {
		// handle error
	}
	return g
}

func deanimate(w http.ResponseWriter, r *http.Request) {
	gifPath := r.URL.Query().Get("url")
	g := fetchGif(gifPath)
	gif.Encode(w, g.Image[0], nil)

}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok")
}

func main() {
	http.HandleFunc("/deanimate", deanimate)
	http.HandleFunc("/healthcheck", healthcheck)
	http.ListenAndServe(":5000", nil)
}
