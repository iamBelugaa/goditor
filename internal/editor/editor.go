package editor

import "github.com/iamBelugaa/editor/internal/command"

// State represents the complete state of the text editor at a point in time.
type State struct {
	Text    string          // The actual text content.
	Command command.Command // The command that led to this state (nil for initial state).
}

// Editor implements a text editor with undo/redo functionality.
type Editor struct {
	history    []*State // Array of all states (acts as our "timeline").
	maxHistory uint     // Maximum number of states to keep in memory.
	currPos    int      // Current position in the history (-1 means no history).
}

// NewTextEditor creates a new text editor instance.
func NewTextEditor(maxHistory uint) *Editor {
	if maxHistory == 0 {
		maxHistory = 100
	}

	return &Editor{
		currPos:    -1,
		maxHistory: maxHistory,
		history:    make([]*State, 0, maxHistory+1),
	}
}

// Text returns the current text content.
func (e *Editor) Text() string {
	if e.currPos == -1 {
		return ""
	}
	return e.history[e.currPos].Text
}

// CanUndo checks if undo operation is possible.
func (e *Editor) CanUndo() bool {
	return e.currPos >= 0
}

// Undo reverts to the previous state, returns true if successful.
func (e *Editor) Undo() bool {
	if !e.CanUndo() {
		return false
	}
	e.currPos--
	return true
}

// CanRedo checks if redo operation is possible.
func (e *Editor) CanRedo() bool {
	return e.currPos < len(e.history)-1
}

// Redo moves forward to the next state, returns true if successful.
func (e *Editor) Redo() bool {
	if !e.CanRedo() {
		return false
	}
	e.currPos++
	return true
}
