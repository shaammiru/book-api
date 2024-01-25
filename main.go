package main

import (
	"fmt"
	"net/http"
)

func listBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of Books!")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Just a Book!")
}

func main() {
	http.HandleFunc("/books", listBooks)
	http.HandleFunc("/books/{id}", getBook)

	fmt.Println("Server runnin at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}