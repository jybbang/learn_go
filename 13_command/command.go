package main

type command interface {
	execute()
	undo()
}
