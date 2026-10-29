package component

import (
	"gotui/buffer"
	"gotui/event"
	"gotui/style"
	"time"
)

type Spinner struct {
	Base
	Frame int
	Label string
	St    style.Style
	stop  chan struct{}
}

var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

func NewSpinner(label string) *Spinner {
	s := &Spinner{Label: label, St: style.Style{Fg: style.Cyan}}
	s.Base.Focusable = false
	s.stop = make(chan struct{})
	go func() {
		t := time.NewTicker(80 * time.Millisecond)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				s.Frame++
			case <-s.stop:
				return
			}
		}
	}()
	return s
}

func (s *Spinner) Stop() { close(s.stop) }

func (s *Spinner) Render(buf *buffer.Buffer) {
	r := s.Rect
	frame := spinnerFrames[s.Frame%len(spinnerFrames)]
	buf.WriteString(r.X, r.Y, frame+" "+s.Label, s.St)
}

func (s *Spinner) Handle(event.Event) bool { return false }
