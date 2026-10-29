package component

import (
	"gotui/buffer"
	"gotui/event"
	"gotui/style"
)

type Rect struct {
	X, Y, W, H int
}

type Component interface {
	Bounds() Rect
	SetBounds(r Rect)
	Render(b *buffer.Buffer)
	Handle(e event.Event) bool
	Children() []Component
	Add(c Component)
	Flex() int
	SetFlex(int)
	PrefH() int
	PrefW() int
	SetPrefH(int)
	SetPrefW(int)
}

type Base struct {
	Rect       Rect
	Style      style.Style
	Kids       []Component
	Focusable  bool
	FlexGrow   int
	PrefHeight int
	PrefWidth  int
}

func (b *Base) Bounds() Rect            { return b.Rect }
func (b *Base) SetBounds(r Rect)        { b.Rect = r }
func (b *Base) Children() []Component   { return b.Kids }
func (b *Base) Add(c Component)         { b.Kids = append(b.Kids, c) }
func (b *Base) Handle(event.Event) bool { return false }
func (b *Base) Flex() int               { return b.FlexGrow }
func (b *Base) SetFlex(f int)           { b.FlexGrow = f }
func (b *Base) PrefH() int              { return b.PrefHeight }
func (b *Base) PrefW() int              { return b.PrefWidth }
func (b *Base) SetPrefH(h int)          { b.PrefHeight = h }
func (b *Base) SetPrefW(w int)          { b.PrefWidth = w }

func CollectFocusable(c Component, out *[]Component) {
	if c == nil {
		return
	}
	if fr, ok := c.(interface{ FocusableFlag() bool }); ok && fr.FocusableFlag() {
		*out = append(*out, c)
	}
	// fallback: check Base via interface
	for _, ch := range c.Children() {
		CollectFocusable(ch, out)
	}
}

func (b *Base) FocusableFlag() bool { return b.Focusable }
