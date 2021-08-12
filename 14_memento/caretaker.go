package main

import "fmt"

type iCaretaker interface {
	execute(iCommand)
	undo()
	canUndo() bool
}

type caretaker struct {
	history []iCommand
}

func (c *caretaker) execute(cmd iCommand) {
	cmd.execute()
	c.history = append(c.history, cmd)
}

func (c *caretaker) undo() {
	if !c.canUndo() {
		fmt.Println("undo unavailable")
		return
	}

	cmd := c.history[len(c.history)-1]
	c.history = c.history[:len(c.history)-1]

	fmt.Print("undo - ")
	cmd.undo()
}

func (c *caretaker) canUndo() bool {
	return len(c.history) > 0
}
