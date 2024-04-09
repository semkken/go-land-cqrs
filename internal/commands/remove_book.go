package commands

type RemoveBookCommand struct {
	BookID string
}

func (cmd *RemoveBookCommand) Execute() error {
	return nil
}
