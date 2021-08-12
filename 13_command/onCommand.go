package main

// command interface
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

func (c *onCommand) undo() {
	c.device.off()
}
