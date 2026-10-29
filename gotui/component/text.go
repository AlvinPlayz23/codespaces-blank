package component

import (
	"gotui/buffer"
	"gotui/event"
	"gotui/style"
	"strings"
)

type Text struct {
	Base
	Content string
	St      style.Style
	Wrap    bool
}

func NewText(content string, st style.Style) *Text {
	t := &Text{Content: content, St: st}
	t.SetPrefH(1 + strings.Count(content, "\n"))
	t.Base.Focusable = false
	return t
}

func (t *Text) Render(buf *buffer.Buffer) {
	r := t.Rect
	lines := strings.Split(t.Content, "\n")
	y := r.Y
	for i, line := range lines {
		if y >= r.Y+r.H {
			break
		}
		buf.WriteString(r.X, y, line, t.St)
		y++
		_ = i
	}
}

func (t *Text) Handle(event.Event) bool { return false }
