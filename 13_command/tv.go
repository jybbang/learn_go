package main

import "fmt"

// device interface
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	if t.isRunning {
		fmt.Println("Already tv on")
		return
	}
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	if !t.isRunning {
		fmt.Println("Already tv off")
		return
	}
	t.isRunning = false
	fmt.Println("Turning tv off")
}
