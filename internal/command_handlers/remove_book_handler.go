package command_handlers

import (
	"net/http"

	"internal/commands"

	"github.com/gorilla/mux"
)

func RemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	bookID := mux.Vars(r)["id"]

	cmd := &commands.RemoveBookCommand{
		BookID: bookID,
	}

	if err := cmd.Execute(); err != nil {
		http.Error(w, "Failed to remove book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
