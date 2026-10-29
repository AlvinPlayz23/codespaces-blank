package buffer

import "gotui/style"

type Cell struct {
	Rune  rune
	Style style.Style
}

type Buffer struct {
	W, H  int
	Cells [][]Cell
}

func New(w, h int) *Buffer {
	b := &Buffer{W: w, H: h}
	b.Cells = make([][]Cell, h)
	for y := 0; y < h; y++ {
		b.Cells[y] = make([]Cell, w)
		for x := 0; x < w; x++ {
			b.Cells[y][x] = Cell{Rune: ' '}
		}
	}
	return b
}

func (b *Buffer) Resize(w, h int) {
	b.W, b.H = w, h
	b.Cells = make([][]Cell, h)
	for y := 0; y < h; y++ {
		b.Cells[y] = make([]Cell, w)
		for x := 0; x < w; x++ {
			b.Cells[y][x] = Cell{Rune: ' '}
		}
	}
}

func (b *Buffer) Set(x, y int, r rune, s style.Style) {
	if x < 0 || y < 0 || x >= b.W || y >= b.H {
		return
	}
	b.Cells[y][x] = Cell{Rune: r, Style: s}
}

func (b *Buffer) WriteString(x, y int, s string, st style.Style) {
	cx := x
	for _, r := range s {
		b.Set(cx, y, r, st)
		cx++
	}
}

func (b *Buffer) FillRect(x, y, w, h int, r rune, s style.Style) {
	for yy := y; yy < y+h; yy++ {
		for xx := x; xx < x+w; xx++ {
			b.Set(xx, yy, r, s)
		}
	}
}

// Render to ANSI string (full redraw)
func (b *Buffer) Render() string {
	var out string
	out += "\x1b[H" // home
	lastStyle := ""
	for y := 0; y < b.H; y++ {
		for x := 0; x < b.W; x++ {
			c := b.Cells[y][x]
			ansi := c.Style.ANSI()
			if ansi != lastStyle {
				out += "\x1b[0m" + ansi
				lastStyle = ansi
			}
			out += string(c.Rune)
		}
		if y < b.H-1 {
			out += "\r\n"
		}
	}
	out += style.Reset
	return out
}
