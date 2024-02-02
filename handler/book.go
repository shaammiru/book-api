package handler

import (
	"net/http"

	"github.com/syamilh/book-api/data"
	"github.com/syamilh/book-api/helper"
)

type BookHandler struct{}

var bookData = &data.BookData{}

func (b *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookData.Create(data.Book{
		ID:       3,
		Title:    "The Lean Startup",
		Author:   "Eric Ries",
		Year:     2011,
		Language: "English",
		Pages:    296,
	})

	helper.WriteJSONMessage(w, http.StatusCreated, "Book created")
}

func (b *BookHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books := bookData.List()

	jsonResponse, err := helper.ToJSON(books)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusInternalServerError, "Error converting data to JSON")
		return
	}

	helper.WriteJSONData(w, http.StatusOK, jsonResponse)
}

func (b *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	i, err := helper.GetIDFromURL(r)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	book, err := bookData.GetByID(i)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusNotFound, "Book not found")
		return
	}

	jsonResponse, err := helper.ToJSON(book)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusInternalServerError, "Error converting data to JSON")
		return
	}

	helper.WriteJSONData(w, http.StatusOK, jsonResponse)
}

func (b *BookHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	i, err := helper.GetIDFromURL(r)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = bookData.UpdateByID(i, data.Book{
		ID:       1,
		Title:    "Harry Potter",
		Author:   "J.K. Rowling",
		Year:     1997,
		Language: "English",
		Pages:    223,
	})
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusNotFound, "Book not found")
		return
	}

	helper.WriteJSONMessage(w, http.StatusOK, "Book updated")
}

func (b *BookHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	i, err := helper.GetIDFromURL(r)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = bookData.DeleteByID(i)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusNotFound, "Book not found")
		return
	}

	helper.WriteJSONMessage(w, http.StatusOK, "Book deleted")
}
