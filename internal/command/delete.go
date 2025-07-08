package command

// DeleteCommand represents deleting text from a specific range.
type DeleteCommand struct {
	Start uint // Starting position of deletion (inclusive).
	End   uint // Ending position of deletion (exclusive).
}

// Execute performs the delete operation.
func (cmd *DeleteCommand) Execute(text string) string {
	runes := []rune(text)

	// Validate and adjust positions.
	if cmd.End > uint(len(runes)) {
		cmd.End = uint(len(runes))
	}

	// Nothing to delete.
	if cmd.Start >= cmd.End {
		return text
	}

	before := runes[:cmd.Start]
	after := runes[cmd.End:]

	// Create new slice without the deleted portion.
	buffer := make([]rune, len(before)+len(after))
	copy(buffer, before)
	copy(buffer[len(before):], after)

	return string(buffer)
}
