package command_handlers

import (
	"encoding/json"
	"internal/commands"
	"net/http"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var bookData struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		ISBN   string `json:"isbn"`
	}

	if err := json.NewDecoder(r.Body).Decode(&bookData); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	cmd := &commands.AddBookCommand{
		Book: &internal.Book{
			Title:  bookData.Title,
			Author: bookData.Author,
			ISBN:   bookData.ISBN,
		},
	}

	if err := cmd.Execute(); err != nil {
		http.Error(w, "Failed to add book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
