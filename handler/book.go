package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/syamilh/book-api/data"
	"github.com/syamilh/book-api/helper"
)

type BookHandler struct{}

var bookData = &data.BookData{}
var validate = validator.New()

func (b *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book data.Book
	if err := helper.FromJSON(r, &book); err != nil {
		helper.WriteJSONMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validate.Struct(book); err != nil {
		validationError := err.(validator.ValidationErrors)
		helper.WriteJSONMessage(w, http.StatusBadRequest, validationError[0].Error())
		return
	}

	if err := bookData.Create(book); err != nil {
		helper.WriteJSONMessage(w, http.StatusConflict, err.Error())
		return
	}

	helper.WriteJSONMessage(w, http.StatusCreated, "Book created")
}

func (b *BookHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books := bookData.List()

	jsonResponse, err := helper.ToJSON(books)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSONData(w, http.StatusOK, jsonResponse)
}

func (b *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	i, err := helper.GetIDFromURL(r)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusBadRequest, "Book not found")
		return
	}

	book, err := bookData.GetByID(i)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusNotFound, err.Error())
		return
	}

	jsonResponse, err := helper.ToJSON(book)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSONData(w, http.StatusOK, jsonResponse)
}

func (b *BookHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book data.Book
	i, err := helper.GetIDFromURL(r)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusBadRequest, "Book not found")
		return
	}

	if err := helper.FromJSON(r, &book); err != nil {
		helper.WriteJSONMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validate.Struct(book); err != nil {
		validationError := err.(validator.ValidationErrors)
		helper.WriteJSONMessage(w, http.StatusBadRequest, validationError[0].Error())
		return
	}

	if err := bookData.UpdateByID(i, book); err != nil {
		helper.WriteJSONMessage(w, http.StatusNotFound, err.Error())
		return
	}

	helper.WriteJSONMessage(w, http.StatusOK, "Book updated")
}

func (b *BookHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	i, err := helper.GetIDFromURL(r)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusBadRequest, "Book not found")
		return
	}

	err = bookData.DeleteByID(i)
	if err != nil {
		helper.WriteJSONMessage(w, http.StatusNotFound, err.Error())
		return
	}

	helper.WriteJSONMessage(w, http.StatusOK, "Book deleted")
}
