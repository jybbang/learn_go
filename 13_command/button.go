package main

type button struct {
	command        command
	commandHistory iCommandHistory
}

func (b *button) press() {
	b.command.execute()
	b.commandHistory.add(b.command)
}
