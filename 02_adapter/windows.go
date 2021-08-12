package main

import "fmt"

type windows struct{}

func (w *windows) insertIntoUSBport() {
	fmt.Println("USB connect is plugged into windows")
}
