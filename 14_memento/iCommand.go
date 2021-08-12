package main

type iCommand interface {
	execute()
	undo()
}
