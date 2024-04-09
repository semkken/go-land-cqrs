package commands

import "internal"

type UpdateBookCommand struct {
	Book *internal.Book
}

func (cmd *UpdateBookCommand) Execute() error {
	return nil
}
