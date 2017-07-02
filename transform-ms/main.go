package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nestorsokil/goproc-img/transform-ms/server"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Running...")
	})

	http.HandleFunc("/gray", server.GrayScaleHandler)
	http.HandleFunc("/binary", server.BinaryHandler)
	http.HandleFunc("/negative", server.NegativeHandler)

	http.HandleFunc("/resize", server.ResizeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
