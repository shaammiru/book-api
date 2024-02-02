package data

import (
	"errors"
)

type Book struct {
	ID       int    `json:"id" validate:"required,number,gt=0"`
	Title    string `json:"title" validate:"required,min=1"`
	Author   string `json:"author" validate:"required,min=1"`
	Year     int    `json:"year" validate:"required,number,gt=0"`
	Language string `json:"language" validate:"required,min=1"`
	Pages    int    `json:"pages" validate:"required,number,gt=0"`
}

type BookData struct{}

var Books = []Book{
	{
		ID:       1,
		Title:    "The Hobbit",
		Author:   "J.R.R. Tolkien",
		Year:     1937,
		Language: "English",
		Pages:    310,
	},
	{
		ID:       2,
		Title:    "Atomic Habits",
		Author:   "James Clear",
		Year:     2018,
		Language: "English",
		Pages:    319,
	},
}

func (bd *BookData) Create(data Book) error {
	for _, book := range Books {
		if book.ID == data.ID {
			return errors.New("Book ID already exists")
		}
	}

	Books = append(Books, data)
	return nil
}

func (bd *BookData) List() []Book {
	return Books
}

func (bd *BookData) GetByID(id int) (Book, error) {
	for _, book := range Books {
		if book.ID == id {
			return book, nil
		}
	}

	return Book{}, errors.New("Book not found")
}

func (b *BookData) UpdateByID(id int, data Book) error {
	for i, book := range Books {
		if book.ID == id {
			Books[i] = data
			return nil
		}
	}

	return errors.New("Book not found")
}

func (b *BookData) DeleteByID(id int) error {
	for i, book := range Books {
		if book.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			return nil
		}
	}

	return errors.New("Book not found")
}
