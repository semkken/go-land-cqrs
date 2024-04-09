package internal

type Repository interface {
	AddBook(book *Book) error
	UpdateBook(book *Book) error
	RemoveBook(id string) error
	GetBookByID(id string) (*Book, error)
	GetAllBooks() ([]*Book, error)
}
