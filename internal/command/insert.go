package command

// InsertCommand represents inserting text at a specific position.
type InsertCommand struct {
	Position uint
	Text     string
}

// Execute performs the insert operation.
func (cmd *InsertCommand) Execute(text string) string {
	// Convert string to rune slice to handle Unicode properly.
	runes := []rune(text)

	// Ensure position is within bounds.
	if cmd.Position > uint(len(runes)) {
		cmd.Position = uint(len(runes))
	}

	// Split the text at the insertion point and insert new text.
	before := runes[:cmd.Position]
	after := runes[cmd.Position:]
	newText := []rune(cmd.Text)

	// Combine: before + newText + after
	buffer := make([]rune, len(runes)+len(newText))
	copy(buffer, before)
	copy(buffer[len(before):], newText)
	copy(buffer[len(before)+len(newText):], after)

	return string(buffer)
}
