package aggregates

import (
	"errors"
	"internal"
)

type BookAggregate struct {
	Books map[string]*internal.Book
}

func NewBookAggregate() *BookAggregate {
	return &BookAggregate{
		Books: make(map[string]*internal.Book),
	}
}

func (ba *BookAggregate) AddBook(book *internal.Book) error {
	if _, exists := ba.Books[book.ID]; exists {
		return errors.New("book with the same ID already exists")
	}
	ba.Books[book.ID] = book
	return nil
}

func (ba *BookAggregate) UpdateBook(book *internal.Book) error {
	if _, exists := ba.Books[book.ID]; !exists {
		return errors.New("book not found")
	}
	ba.Books[book.ID] = book
	return nil
}

func (ba *BookAggregate) RemoveBook(bookID string) error {
	if _, exists := ba.Books[bookID]; !exists {
		return errors.New("book not found")
	}
	delete(ba.Books, bookID)
	return nil
}

func (ba *BookAggregate) GetBookByID(bookID string) (*internal.Book, error) {
	book, exists := ba.Books[bookID]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

func (ba *BookAggregate) GetAllBooks() ([]*internal.Book, error) {
	books := make([]*internal.Book, 0, len(ba.Books))
	for _, book := range ba.Books {
		books = append(books, book)
	}
	return books, nil
}
