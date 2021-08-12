package main

type Size uint8

const (
	small Size = iota
	medium
	large
	extraLarge
)

type BodyPart int

const (
	Head     BodyPart = iota // Head = 0
	Shoulder                 // Shoulder = 1
	Knee                     // Knee = 2
	Toe                      // Toe = 3
)

func (bp BodyPart) String() string {
	return []string{"Head", "Shoulder", "Knee", "Toe"}[bp]
}
