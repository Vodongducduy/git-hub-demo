package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		return
	}
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	log.Println("Hello world")
}
