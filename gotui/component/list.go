package component

import (
	"gotui/buffer"
	"gotui/event"
	"gotui/style"
)

type ListItem struct {
	Label string
	Value string
}

type List struct {
	Base
	Items    []ListItem
	Selected int
	St       style.Style
	SelSt    style.Style
	OnSelect func(int)
}

func NewList(items []ListItem) *List {
	l := &List{Items: items, St: style.Style{Fg: style.White}, SelSt: style.Style{Fg: style.Black, Bg: style.Cyan, Bold: true}}
	l.Base.Focusable = true
	return l
}

func (l *List) Render(buf *buffer.Buffer) {
	r := l.Rect
	for i, it := range l.Items {
		if i >= r.H {
			break
		}
		st := l.St
		if i == l.Selected {
			st = l.SelSt
		}
		// fill line
		for x := r.X; x < r.X+r.W; x++ {
			buf.Set(x, r.Y+i, ' ', st)
		}
		buf.WriteString(r.X+1, r.Y+i, it.Label, st)
	}
}

func (l *List) Handle(e event.Event) bool {
	if e.Type != event.EvKey {
		return false
	}
	switch e.Key.Name {
	case "up":
		if l.Selected > 0 {
			l.Selected--
		}
		return true
	case "down":
		if l.Selected < len(l.Items)-1 {
			l.Selected++
		}
		return true
	case "enter":
		if l.OnSelect != nil {
			l.OnSelect(l.Selected)
		}
		return true
	}
	return false
}
