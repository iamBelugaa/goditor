# Text Editor with Undo/Redo Using the Command Pattern

This project demonstrates how to build a text editor that can undo and redo
operations using the Command pattern.

## Understanding the Command Pattern

The Command pattern works by wrapping each action in an object that knows how to
perform that action. Instead of calling methods directly on your data, you
create command objects that contain all the information needed to perform an
operation.

Think of it like writing instructions on index cards. Instead of immediately
doing something when someone asks, you write down exactly what they want on a
card, then follow the instructions on the card. This way, you have a stack of
instruction cards that you can go through one by one, skip some, or even go
backwards through them.

In our text editor, every time someone wants to insert text or delete text, we
don't change the text directly. Instead, we create a command object that knows
exactly what change to make. Then we tell that command to execute itself. After
it runs, we save both the command and the result, so we can see what happened
and potentially go back to how things were before.

## How the Internal Mechanics Work

When you create a new text editor, it starts with an empty history and a current
position of negative one, which means there's no text yet. The history is like a
timeline where each point represents what the text looked like after some
operation.

Let's trace through what happens step by step when you perform operations. When
you call `Insert(0, "Hello")`, the editor first creates an InsertCommand object.
This object stores the position where text should be inserted (position 0) and
the text to insert ("Hello"). The editor then asks this command to execute
itself by calling `Execute()` on the current text, which starts as an empty
string.

The InsertCommand's Execute method converts the current text into an array of
Unicode characters (called runes in Go), finds the insertion point, and creates
a new string with the new text inserted at that position. Since we're inserting
at position 0 in an empty string, the result is simply "Hello".

The editor takes this result and creates a new State object that contains both
the resulting text ("Hello") and the command that produced it. This state gets
added to the history array, and the current position moves to point at this new
state.

Now when you call `Insert(5, " World")`, the same process happens, but this time
the InsertCommand executes against the current text "Hello". The command finds
position 5 (which is at the end of "Hello"), inserts " World", and returns
"Hello World". Again, this creates a new state in the history.

The undo mechanism works by simply moving backwards through this history
timeline. When you call `Undo()`, the editor decreases the current position by
one, which makes the previous state become the current one. The editor doesn't
need to figure out how to reverse the operation because it already has the text
from before the operation was performed.

Redo works the opposite way. If you've undone some operations, there are states
in the history that come after your current position. Calling `Redo()` moves the
current position forward to the next state in the history.

## Example: Building a Shopping List

Imagine we're using our text editor to build a shopping list.

We start by creating an editor with a history limit of 10 operations:

```go
editor := editor.NewTextEditor(10)
```

At this point, the editor has an empty history array and a current position of
-1, meaning no operations have been performed yet.

First, we add the title of our list:

```go
editor.Insert(0, "Shopping List:\n")
```

The editor creates an InsertCommand with position 0 and text "Shopping List:\n".
Since there's no current text (empty string), the command's Execute method
returns "Shopping List:\n". The editor creates a new state containing this text
and the command, adds it to the history, and sets the current position to 0.

Next, we add our first item:

```go
editor.Insert(15, "- Milk\n")
```

Now the InsertCommand executes against the current text "Shopping List:\n". It
finds position 15 (right at the end), inserts "- Milk\n", and returns "Shopping
List:\n- Milk\n". This becomes state 1 in our history, and the current position
moves to 1.

We continue adding items:

```go
editor.Insert(23, "- Bread\n")  // Results in "Shopping List:\n- Milk\n- Bread\n"
editor.Insert(32, "- Eggs\n")   // Results in "Shopping List:\n- Milk\n- Bread\n- Eggs\n"
```

At this point, our history contains four states (positions 0 through 3), each
representing what the shopping list looked like after each insertion.

Now suppose we realize we don't need eggs. We can delete them:

```go
editor.Delete(32, 39)  // Delete "- Eggs\n"
```

The editor creates a DeleteCommand with start position 32 and end position 39.
This command's Execute method removes characters from position 32 to 38
(remember, the end position is exclusive), returning "Shopping List:\n- Milk\n-
Bread\n". This becomes state 4 in our history.

If we change our mind about removing eggs, we can undo:

```go
editor.Undo()  // Goes back to state 3: "Shopping List:\n- Milk\n- Bread\n- Eggs\n"
```

The editor simply decreases the current position from 4 to 3, making the
previous state current again. No computation is needed because we already saved
what the text looked like at that point.

If we undo again:

```go
editor.Undo()  // Goes back to state 2: "Shopping List:\n- Milk\n- Bread\n"
```

Now if we decide to add a different item instead of eggs:

```go
editor.Insert(32, "- Cheese\n")
```

Something important happens here. Because we're not at the end of our history
when we perform this new operation, the editor removes all the states that came
after our current position. States 3 and 4 (which contained the versions with
eggs) get discarded, and the new state with cheese becomes the new end of our
history.

This behavior ensures that history remains linear and prevents confusing
situations where you could redo to multiple different futures.
