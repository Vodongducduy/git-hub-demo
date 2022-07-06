package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/googlebye", goodBye)
	http.ListenAndServe(":9090", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world")
	d, err := ioutil.ReadAll((r.Body))
	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s", d)
}

func goodBye(w http.ResponseWriter, r *http.Request) {
	log.Println("Good bye")
}
