package main

import (
	"fmt"
	"github.com/a-h/templ"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	s := Hello("Nice")
	http.HandleFunc("/add", add)
	http.HandleFunc("/", templ.Handler(s))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func add(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Add")
}
