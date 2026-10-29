package gotui

import (
	"fmt"
	"os"
	"time"

	"gotui/buffer"
	"gotui/component"
	"gotui/event"
	"gotui/term"
)

type App struct {
	Term    *term.Terminal
	Root    *component.Box
	Focused component.Component
	Events  chan event.Event
	Quit    chan struct{}
	Dirty   bool
}

func NewApp(root *component.Box) *App {
	return &App{
		Term:   term.New(),
		Root:   root,
		Events: make(chan event.Event, 256),
		Quit:   make(chan struct{}),
		Dirty:  true,
	}
}

func (a *App) Run() error {
	if err := a.Term.EnterRaw(); err != nil {
		return err
	}
	defer a.Term.ExitRaw()

	// input reader goroutine
	go a.readInput()

	// resize ticker
	go func() {
		t := time.NewTicker(250 * time.Millisecond)
		defer t.Stop()
		lastW, lastH := a.Term.Size()
		for {
			select {
			case <-a.Quit:
				return
			case <-t.C:
				w, h := a.Term.Size()
				if w != lastW || h != lastH {
					lastW, lastH = w, h
					a.Events <- event.Event{Type: event.Resize, W: w, H: h}
				}
			}
		}
	}()

	// initial
	w, h := a.Term.Size()
	a.Root.SetBounds(component.Rect{X: 0, Y: 0, W: w, H: h})

	for {
		select {
		case <-a.Quit:
			return nil
		case e := <-a.Events:
			switch e.Type {
			case event.Quit:
				return nil
			case event.Resize:
				a.Root.SetBounds(component.Rect{X: 0, Y: 0, W: e.W, H: e.H})
				a.Dirty = true
			case event.EvKey:
				if e.Key.Name == "tab" {
					// cycle focus
					var focuses []component.Component
					component.CollectFocusable(a.Root, &focuses)
					if len(focuses) > 0 {
						idx := -1
						for i, c := range focuses {
							if c == a.Focused {
								idx = i
							}
						}
						idx = (idx + 1) % len(focuses)
						// blur old
						if inp, ok := a.Focused.(*component.Input); ok {
							inp.Focused = false
						}
						if lst, ok := a.Focused.(*component.List); ok {
							_ = lst
						}
						a.Focused = focuses[idx]
						if inp, ok := a.Focused.(*component.Input); ok {
							inp.Focused = true
						}
					}
					a.Dirty = true
					break
				}
				if e.Key.Name == "ctrl+c" {
					return nil
				}
				if a.Focused != nil {
					a.Focused.Handle(e)
				} else {
					a.Root.Handle(e)
				}
				a.Dirty = true
			}
		default:
			// render if dirty
			if a.Dirty {
				w, h = a.Term.Size()
				if a.Root.Bounds().W != w || a.Root.Bounds().H != h {
					a.Root.SetBounds(component.Rect{X: 0, Y: 0, W: w, H: h})
				}
				buf := buffer.New(w, h)
				a.Root.Render(buf)
				fmt.Print(buf.Render())
				a.Dirty = false
			}
			time.Sleep(16 * time.Millisecond)
		}
	}
}

func (a *App) readInput() {
	buf := make([]byte, 256)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			a.Events <- event.Event{Type: event.Quit}
			return
		}
		if n == 0 {
			continue
		}
		i := 0
		for i < n {
			b := buf[i]
			// escape sequences
			if b == 0x1b {
				if i+1 < n && buf[i+1] == '[' {
					// CSI
					j := i + 2
					for j < n && (buf[j] < 0x40 || buf[j] > 0x7e) {
						j++
					}
					if j < n {
						code := buf[j]
						switch code {
						case 'A':
							a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Name: "up"}}
						case 'B':
							a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Name: "down"}}
						case 'C':
							a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Name: "right"}}
						case 'D':
							a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Name: "left"}}
						case 'H':
							a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Name: "home"}}
						case 'F':
							a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Name: "end"}}
						}
						i = j + 1
						continue
					}
				}
				// lone esc
				a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Name: "esc"}}
				i++
				continue
			}
			switch b {
			case 0x03: // ctrl+c
				a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Name: "ctrl+c", Ctrl: true}}
			case 0x0d, 0x0a: // enter
				a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Name: "enter"}}
			case 0x7f, 0x08: // backspace
				a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Name: "backspace"}}
			case 0x09: // tab
				a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Name: "tab"}}
			default:
				if b >= 0x20 {
					a.Events <- event.Event{Type: event.EvKey, Key: event.Key{Rune: rune(b)}}
				}
			}
			i++
		}
	}
}
