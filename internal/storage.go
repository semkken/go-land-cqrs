package internal

import "errors"

type MemoryStorage struct {
	books map[string]*Book
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		books: make(map[string]*Book),
	}
}

func (s *MemoryStorage) AddBook(book *Book) error {
	if _, exists := s.books[book.ID]; exists {
		return errors.New("book with the same ID already exists")
	}
	s.books[book.ID] = book
	return nil
}

func (s *MemoryStorage) UpdateBook(book *Book) error {
	if _, exists := s.books[book.ID]; !exists {
		return errors.New("book not found")
	}
	s.books[book.ID] = book
	return nil
}

func (s *MemoryStorage) RemoveBook(id string) error {
	if _, exists := s.books[id]; !exists {
		return errors.New("book not found")
	}
	delete(s.books, id)
	return nil
}

func (s *MemoryStorage) GetBookByID(id string) (*Book, error) {
	book, exists := s.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

func (s *MemoryStorage) GetAllBooks() ([]*Book, error) {
	books := make([]*Book, 0, len(s.books))
	for _, book := range s.books {
		books = append(books, book)
	}
	return books, nil
}
