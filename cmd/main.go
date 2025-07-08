package main

import (
	"fmt"

	"github.com/iamBelugaa/goditor/internal/editor"
)

func main() {
	fmt.Println("=== Text Editor with Undo/Redo Demo ===")

	editor := editor.NewTextEditor(15)

	fmt.Println("\n--- Performing Operations ---")

	editor.Insert(0, "Hello")
	fmt.Printf("After inserting 'Hello': \"%s\"\n", editor.Text())

	editor.Insert(5, " World")
	fmt.Printf("After inserting ' World': \"%s\"\n", editor.Text())

	editor.Insert(11, "!")
	fmt.Printf("After inserting '!': \"%s\"\n", editor.Text())

	editor.Delete(5, 11) // Delete " World"
	fmt.Printf("After deleting ' World': \"%s\"\n", editor.Text())

	editor.Insert(5, " Go")
	fmt.Printf("After inserting ' Go': \"%s\"\n", editor.Text())

	fmt.Println("\n--- Testing Undo Operations ---")

	for editor.CanUndo() {
		fmt.Printf("Before undo: \"%s\"\n", editor.Text())
		editor.Undo()
		fmt.Printf("After undo: \"%s\"\n", editor.Text())
		fmt.Println()
	}

	fmt.Println("Attempting to undo when at beginning...")
	if !editor.Undo() {
		fmt.Println("Cannot undo - already at the beginning of history")
	}

	fmt.Println("\n--- Testing Redo Operations ---")

	for i := 0; i < 3 && editor.CanRedo(); i++ {
		fmt.Printf("Before redo: \"%s\"\n", editor.Text())
		editor.Redo()
		fmt.Printf("After redo: \"%s\"\n", editor.Text())
		fmt.Println()
	}

	fmt.Println("--- Testing History Truncation ---")
	fmt.Println("Current state before new operation:", editor.Text())

	editor.Insert(11, ", Programming")
	fmt.Printf("After inserting ', Programming': \"%s\"\n", editor.Text())

	editor.Insert(editor.Length()-1, " Meow Meow")
	fmt.Printf("After inserting ' Meow Meow': \"%s\"\n", editor.Text())
}
