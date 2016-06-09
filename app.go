package main

import (
	"net/http"
	"image/gif"
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

func main() {
	http.HandleFunc("/deanimate", deanimate)
	http.ListenAndServe(":7896", nil)
}
