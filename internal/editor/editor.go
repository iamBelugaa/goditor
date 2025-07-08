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
	if e.currPos < 0 {
		return ""
	}
	return e.history[e.currPos].Text
}

// Insert adds text at the specified position.
func (e *Editor) Insert(pos uint, text string) {
	if text == "" {
		return
	}

	cmd := &command.InsertCommand{Position: pos, Text: text}
	newText := cmd.Execute(e.Text())

	e.addHistory(newText, cmd)
}

// Delete removes text from the specified range.
func (e *Editor) Delete(start, end uint) {
	if start >= end || start >= uint(len(e.Text())) {
		return
	}

	cmd := &command.DeleteCommand{Start: start, End: end}
	newText := cmd.Execute(e.Text())

	e.addHistory(newText, cmd)
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

// Info returns information about the current state and available operations.
func (e *Editor) Info() (currentPos int, totalHistory int, canUndo bool, canRedo bool) {
	return e.currPos, len(e.history), e.CanUndo(), e.CanRedo()
}

// addHistory adds a new state to the history, managing the circular buffer if needed.
func (e *Editor) addHistory(text string, cmd command.Command) {
	// If we're not at the end of history, we need to truncate future states.
	// This happens when user undoes some operations and then performs a new one.
	if e.currPos < len(e.history)-1 {
		e.history = e.history[:e.currPos+1]
	}

	// Add the new state.
	e.history = append(e.history, &State{Text: text, Command: cmd})
	e.currPos = len(e.history) - 1

	// If we exceed maximum history, remove the oldest state.
	if e.currPos > int(e.maxHistory) {
		// Shift all elements left by one position.
		e.history = e.history[1:]
		e.currPos--
	}
}
