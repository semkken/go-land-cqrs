package commands

import "internal"

type AddBookCommand struct {
	Book *internal.Book
}

func (cmd *AddBookCommand) Execute() error {
	return nil
}
