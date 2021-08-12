package main

import "fmt"

type iCommandHistory interface {
	add(command)
	undo()
	canUndo() bool
}

type commandHistory struct {
	undos []command
}

func (c *commandHistory) add(cmd command) {
	c.undos = append(c.undos, cmd)
}

func (c *commandHistory) undo() {
	if !c.canUndo() {
		fmt.Println("undo unavailable")
		return
	}

	cmd := c.undos[len(c.undos)-1]
	c.undos = c.undos[:len(c.undos)-1]

	fmt.Print("undo - ")
	cmd.undo()
}

func (c *commandHistory) canUndo() bool {
	return len(c.undos) > 0
}
