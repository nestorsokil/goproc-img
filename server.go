package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Running...")
	})

	http.HandleFunc("/gray", GrayScaleHandler)
	http.HandleFunc("/binary", BinaryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
