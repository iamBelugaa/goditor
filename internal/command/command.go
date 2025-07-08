package command

// Command represents an operation that can be performed on the text editor.
type Command interface {
	Execute(text string) string
}
