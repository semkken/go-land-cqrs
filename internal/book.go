package internal

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func NewBook(id, title, author, isbn string) *Book {
	return &Book{
		ID:     id,
		Title:  title,
		Author: author,
		ISBN:   isbn,
	}
}
