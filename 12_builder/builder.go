package main

type builder struct {
	instance house
}

func getBuilder() iBuilder {
	return &builder{}
}

func (b *builder) setWindowType(window string) iBuilder {
	b.instance.windowType = window
	return b
}
func (b *builder) setDoorType(door string) iBuilder {
	b.instance.doorType = door
	return b
}
func (b *builder) setNumFloor(floor int) iBuilder {
	b.instance.floor = floor
	return b
}
func (b *builder) build() *house {
	return &b.instance
}
