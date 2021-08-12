package main

type iBuilder interface {
	setWindowType(string) iBuilder
	setDoorType(string) iBuilder
	setNumFloor(int) iBuilder
	build() *house
}
