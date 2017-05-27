package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Running...")
	})

	http.HandleFunc("/gray", GrayScaleHandler)
	http.HandleFunc("/binary", BinaryHandler)
	http.HandleFunc("/negative", NegativeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
