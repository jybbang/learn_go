package main

import "fmt"

type onCommand struct {
	originator *originator
	memento    *memento
}

func (c *onCommand) execute() {
	c.memento = c.originator.save()
	c.originator.state = "on"
	fmt.Println("on execute")
}

func (c *onCommand) undo() {
	c.originator.restore(c.memento)
	fmt.Println("on undo")
}
