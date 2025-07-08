package editor

import "github.com/iamBelugaa/editor/internal/command"

// State represents the complete state of the text editor at a point in time.
type State struct {
	Text    string          // The actual text content.
	Command command.Command // The command that led to this state (nil for initial state).
}

// Editor implements a text editor with undo/redo functionality.
type Editor struct {
	history []*State // Array of all states (acts as our "timeline").
	currPos int      // Current position in the history (-1 means no history).
}

// NewTextEditor creates a new text editor instance.
func NewTextEditor(maxHistory uint) *Editor {
	if maxHistory == 0 {
		maxHistory = 100
	}

	return &Editor{
		currPos: -1,
		history: make([]*State, 0, maxHistory+1),
	}
}
