package term

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type Terminal struct {
	Fd  int
	Old *term.State
	W   int
	H   int
}

func New() *Terminal {
	fd := int(os.Stdin.Fd())
	t := &Terminal{Fd: fd}
	w, h, err := term.GetSize(fd)
	if err != nil {
		w, h = 80, 24
	}
	t.W, t.H = w, h
	return t
}

func (t *Terminal) EnterRaw() error {
	old, err := term.MakeRaw(t.Fd)
	if err != nil {
		return err
	}
	t.Old = old
	// alternate screen, hide cursor, clear
	fmt.Print("\x1b[?1049h\x1b[?25l\x1b[2J\x1b[H")
	return nil
}

func (t *Terminal) ExitRaw() {
	if t.Old != nil {
		term.Restore(t.Fd, t.Old)
	}
	fmt.Print("\x1b[?25h\x1b[?1049l")
}

func (t *Terminal) Size() (int, int) {
	w, h, err := term.GetSize(t.Fd)
	if err != nil {
		return t.W, t.H
	}
	t.W, t.H = w, h
	return w, h
}
