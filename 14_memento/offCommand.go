package main

import "fmt"

type offCommand struct {
	originator *originator
	memento    *memento
}

func (c *offCommand) execute() {
	c.memento = c.originator.save()
	c.originator.state = "off"
	fmt.Println("off execute")
}

func (c *offCommand) undo() {
	c.originator.restore(c.memento)
	fmt.Println("off undo")
}
