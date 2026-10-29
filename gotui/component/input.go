package component

import (
	"gotui/buffer"
	"gotui/event"
	"gotui/style"
)

type Input struct {
	Base
	Value       string
	Cursor      int
	Focused     bool
	Placeholder string
	St          style.Style
	OnSubmit    func(string)
}

func NewInput() *Input {
	in := &Input{St: style.Style{Fg: style.White}}
	in.SetPrefH(1)
	in.Base.Focusable = true
	return in
}

func (in *Input) Render(buf *buffer.Buffer) {
	r := in.Rect
	// background
	for x := r.X; x < r.X+r.W; x++ {
		buf.Set(x, r.Y, ' ', style.Style{Bg: style.Bg256(236)})
	}
	text := in.Value
	if text == "" && !in.Focused {
		text = in.Placeholder
		buf.WriteString(r.X, r.Y, text, style.Style{Fg: style.Gray})
	} else {
		buf.WriteString(r.X, r.Y, text, in.St)
	}
	if in.Focused {
		cx := r.X + in.Cursor
		if cx < r.X+r.W {
			buf.Set(cx, r.Y, ' ', style.Style{Reverse: true})
		}
	}
}

func (in *Input) Handle(e event.Event) bool {
	if !in.Focused || e.Type != event.EvKey {
		return false
	}
	k := e.Key
	switch k.Name {
	case "enter":
		if in.OnSubmit != nil {
			in.OnSubmit(in.Value)
		}
		return true
	case "backspace":
		if in.Cursor > 0 && len(in.Value) > 0 {
			runes := []rune(in.Value)
			in.Value = string(append(runes[:in.Cursor-1], runes[in.Cursor:]...))
			in.Cursor--
		}
		return true
	case "left":
		if in.Cursor > 0 {
			in.Cursor--
		}
		return true
	case "right":
		if in.Cursor < len([]rune(in.Value)) {
			in.Cursor++
		}
		return true
	case "home":
		in.Cursor = 0
		return true
	case "end":
		in.Cursor = len([]rune(in.Value))
		return true
	default:
		if k.Rune != 0 && k.Name == "" {
			runes := []rune(in.Value)
			// insert
			out := append([]rune{}, runes[:in.Cursor]...)
			out = append(out, k.Rune)
			out = append(out, runes[in.Cursor:]...)
			in.Value = string(out)
			in.Cursor++
			return true
		}
	}
	return false
}
