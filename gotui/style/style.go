package style

import "fmt"

type Color string

const (
	Reset   = "\x1b[0m"
	Bold    = "\x1b[1m"
	Dim     = "\x1b[2m"
	Italic  = "\x1b[3m"
	Under   = "\x1b[4m"
	Reverse = "\x1b[7m"
)

const (
	Black         Color = "\x1b[30m"
	Red           Color = "\x1b[31m"
	Green         Color = "\x1b[32m"
	Yellow        Color = "\x1b[33m"
	Blue          Color = "\x1b[34m"
	Magenta       Color = "\x1b[35m"
	Cyan          Color = "\x1b[36m"
	White         Color = "\x1b[37m"
	Gray          Color = "\x1b[90m"
	BrightRed     Color = "\x1b[91m"
	BrightGreen   Color = "\x1b[92m"
	BrightYellow  Color = "\x1b[93m"
	BrightBlue    Color = "\x1b[94m"
	BrightMagenta Color = "\x1b[95m"
	BrightCyan    Color = "\x1b[96m"
	BrightWhite   Color = "\x1b[97m"
)

func Fg256(n int) Color { return Color(fmt.Sprintf("\x1b[38;5;%dm", n)) }
func Bg256(n int) Color { return Color(fmt.Sprintf("\x1b[48;5;%dm", n)) }
func FgRGB(r, g, b int) Color {
	return Color(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b))
}
func BgRGB(r, g, b int) Color {
	return Color(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b))
}

type Style struct {
	Fg      Color
	Bg      Color
	Bold    bool
	Dim     bool
	Ital    bool
	Reverse bool
}

func (s Style) ANSI() string {
	out := ""
	if s.Reverse {
		out += Reverse
	}
	if s.Bold {
		out += Bold
	}
	if s.Dim {
		out += Dim
	}
	if s.Ital {
		out += Italic
	}
	if s.Fg != "" {
		out += string(s.Fg)
	}
	if s.Bg != "" {
		out += string(s.Bg)
	}
	return out
}
