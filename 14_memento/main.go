package main

import "fmt"

func main() {
	caretaker := &caretaker{}

	originator := &originator{
		state: "none",
	}

	onCommand := &onCommand{
		originator: originator,
	}

	offCommand := &offCommand{
		originator: originator,
	}

	caretaker.execute(onCommand)
	fmt.Println(originator)

	caretaker.execute(offCommand)
	fmt.Println(originator)

	caretaker.undo()
	fmt.Println(originator)

	caretaker.undo()
	fmt.Println(originator)

	caretaker.undo()
	fmt.Println(originator)
}
