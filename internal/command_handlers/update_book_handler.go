package command_handlers

import (
	"encoding/json"
	"net/http"

	"internal/commands"
)

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	var bookData struct {
		ID     string `json:"id"`
		Title  string `json:"title"`
		Author string `json:"author"`
		ISBN   string `json:"isbn"`
	}

	if err := json.NewDecoder(r.Body).Decode(&bookData); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	cmd := &commands.UpdateBookCommand{
		Book: &internal.Book{
			ID:     bookData.ID,
			Title:  bookData.Title,
			Author: bookData.Author,
			ISBN:   bookData.ISBN,
		},
	}

	if err := cmd.Execute(); err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
