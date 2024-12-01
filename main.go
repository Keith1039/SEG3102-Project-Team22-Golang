package main

import (
	"fmt"
	"github.com/Keith1039/SEG3102-Project-Team22-Golang/templates"
	"github.com/a-h/templ"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	s := templates.Hello("Nice")
	http.HandleFunc("/add", add)
	http.Handle("/", templ.Handler(s))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func add(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Add")
}
