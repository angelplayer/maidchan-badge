package main

import (
	"log"
	"net/http"

	"github.com/angelplayer/maidchan-badge/badge"
)

func main() {

	http.Handle("/go", http.HandlerFunc(test))
	err := http.ListenAndServe(":4020", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}

}

func test(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "image/svg+xml")
	badge.Create(res)
}
