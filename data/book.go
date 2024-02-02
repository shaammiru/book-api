package data

import "errors"

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Language string `json:"language"`
	Pages    int    `json:"pages"`
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

func (bd *BookData) Create(data Book) {
	Books = append(Books, data)
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
