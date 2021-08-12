package main

import "fmt"

type windowsAdapter struct {
	windowsMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts lightning to USB")
	w.windowsMachine.insertIntoUSBport()
}
