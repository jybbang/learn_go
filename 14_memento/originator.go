package main

type originator struct {
	state string
}

func (e *originator) clone() *originator {
	return &originator{state: e.state}
}

func (e *originator) save() *memento {
	return &memento{prototype: e.clone()}
}

func (e *originator) restore(m *memento) {
	e.state = m.prototype.state
}
