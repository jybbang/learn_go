package main

func main() {
	tv := &tv{}

	commandHistory := &commandHistory{}

	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command:        onCommand,
		commandHistory: commandHistory,
	}
	onButton.press()

	offButton := &button{
		command:        offCommand,
		commandHistory: commandHistory,
	}
	offButton.press()

	commandHistory.undo()
	commandHistory.undo()
	commandHistory.undo()

	onButton.press()
	commandHistory.undo()
}
