package component

import (
	"gotui/buffer"
	"gotui/event"
	"gotui/style"
)

type Direction int

const (
	Column Direction = iota
	Row
)

type Box struct {
	Base
	Dir         Direction
	PadX, PadY  int
	Border      bool
	Bg          style.Color
	BorderColor style.Color
	Title       string
}

func NewBox(dir Direction) *Box {
	b := &Box{Dir: dir}
	return b
}

func (b *Box) Render(buf *buffer.Buffer) {
	r := b.Rect
	st := style.Style{Bg: b.Bg}
	if b.Bg != "" {
		buf.FillRect(r.X, r.Y, r.W, r.H, ' ', st)
	}
	if b.Border {
		bs := style.Style{Fg: b.BorderColor}
		if b.BorderColor == "" {
			bs.Fg = style.White
		}
		buf.Set(r.X, r.Y, '+', bs)
		buf.Set(r.X+r.W-1, r.Y, '+', bs)
		buf.Set(r.X, r.Y+r.H-1, '+', bs)
		buf.Set(r.X+r.W-1, r.Y+r.H-1, '+', bs)
		for x := r.X + 1; x < r.X+r.W-1; x++ {
			buf.Set(x, r.Y, '-', bs)
			buf.Set(x, r.Y+r.H-1, '-', bs)
		}
		for y := r.Y + 1; y < r.Y+r.H-1; y++ {
			buf.Set(r.X, y, '|', bs)
			buf.Set(r.X+r.W-1, y, '|', bs)
		}
		if b.Title != "" {
			buf.WriteString(r.X+2, r.Y, " "+b.Title+" ", style.Style{Fg: style.Bold, Bg: b.Bg})
		}
	}
	innerX := r.X + b.PadX
	innerY := r.Y + b.PadY
	innerW := r.W - 2*b.PadX
	innerH := r.H - 2*b.PadY
	if b.Border {
		innerX += 1
		innerY += 1
		innerW -= 2
		innerH -= 2
	}
	if innerW < 0 {
		innerW = 0
	}
	if innerH < 0 {
		innerH = 0
	}

	if b.Dir == Column {
		fixed := 0
		flexTotal := 0
		for _, c := range b.Children() {
			if c.Flex() > 0 {
				flexTotal += c.Flex()
			} else {
				h := c.PrefH()
				if h <= 0 {
					h = c.Bounds().H
				}
				if h <= 0 {
					h = 1
				}
				fixed += h
			}
		}
		remaining := innerH - fixed
		if remaining < 0 {
			remaining = 0
		}
		cy := innerY
		for _, c := range b.Children() {
			var h int
			if c.Flex() > 0 {
				if flexTotal > 0 {
					h = remaining * c.Flex() / flexTotal
				}
			} else {
				h = c.PrefH()
				if h <= 0 {
					h = c.Bounds().H
				}
				if h <= 0 {
					h = 1
				}
			}
			if h < 0 {
				h = 0
			}
			c.SetBounds(Rect{X: innerX, Y: cy, W: innerW, H: h})
			c.Render(buf)
			cy += h
		}
	} else {
		fixed := 0
		flexTotal := 0
		for _, c := range b.Children() {
			if c.Flex() > 0 {
				flexTotal += c.Flex()
			} else {
				w := c.PrefW()
				if w <= 0 {
					w = c.Bounds().W
				}
				if w <= 0 {
					w = 1
				}
				fixed += w
			}
		}
		remaining := innerW - fixed
		if remaining < 0 {
			remaining = 0
		}
		cx := innerX
		for _, c := range b.Children() {
			var w int
			if c.Flex() > 0 {
				if flexTotal > 0 {
					w = remaining * c.Flex() / flexTotal
				}
			} else {
				w = c.PrefW()
				if w <= 0 {
					w = c.Bounds().W
				}
				if w <= 0 {
					w = 1
				}
			}
			if w < 0 {
				w = 0
			}
			c.SetBounds(Rect{X: cx, Y: innerY, W: w, H: innerH})
			c.Render(buf)
			cx += w
		}
	}
}

func (b *Box) Handle(e event.Event) bool {
	for _, c := range b.Children() {
		if c.Handle(e) {
			return true
		}
	}
	return false
}
