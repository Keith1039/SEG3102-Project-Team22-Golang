package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/add", add)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func add(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Add")
}
