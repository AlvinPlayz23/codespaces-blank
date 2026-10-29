package event

type Type int

const (
	EvKey Type = iota
	Resize
	Quit
)

type Key struct {
	Rune rune
	Name string
	Ctrl bool
}

type Event struct {
	Type Type
	Key  Key
	W, H int
}
