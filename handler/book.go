package handler

import (
	"fmt"
	"net/http"
)

type Book struct{}

func (b *Book) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create book")
}

func (b *Book) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List book")
}

func (b *Book) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get book")
}

func (b *Book) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update book")
}

func (b *Book) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete book")
}
