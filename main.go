package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/example/goproc-img/server"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Running...")
	})

	http.HandleFunc("/gray", server.GrayScaleHandler)
	http.HandleFunc("/binary", server.BinaryHandler)
	http.HandleFunc("/negative", server.NegativeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}