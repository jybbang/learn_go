package main

// command interface
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

func (c *offCommand) undo() {
	c.device.on()
}
