package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/home/pra/project/wasm")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
