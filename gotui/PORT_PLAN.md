# gotui: Porting OpenTUI to Go

## Goal
Port the OpenTUI terminal UI framework (Zig core + TypeScript bindings) to a
pure-Go library `gotui`, then use it to build a mock of OpenCode's TUI.

## Source layout (cloned)
- `opentui/` — Zig native core (`packages/native/src/*.zig`) + TS core (`packages/core/src`)
- `opencode/` — consumer TUI (`packages/tui/src`), SolidJS over OpenTUI

## Architecture mapping

| OpenTUI (TS/Zig)            | gotui (Go) target                          |
|-----------------------------|--------------------------------------------|
| `Renderer` / `CliRenderer`  | `gotui.App` (terminal lifecycle + loop)     |
| Renderable tree             | `component.Component` interface + `Box`     |
| `OptimizedBuffer`           | `buffer.Buffer` (cell grid + ANSI emit)    |
| Yoga flexbox layout         | `Box` flex grow/pref sizes (phase 2: full Yoga via CGO or pure Go) |
| Input parser / key events  | `event.Event` + stdin CSI parser           |
| Text buffer / editor view  | `text.Buffer` package (phase 3)            |
| Syntax highlighting         | Chroma or tree-sitter Go bindings (phase 4) |
| Native session / VT        | `term.Terminal` + `ghostty-vt` port (phase 5) |

## Phases

1. **Foundation (done)** — terminal raw mode, cell buffer, ANSI render, events, basic components (Box, Text, Input, List, Spinner), flex layout (grow/pref), mock app.
2. **Layout** — full flexbox (Yoga parity): margins, padding, gap, absolute positioning, min/max sizes, text wrapping measurement.
3. **Text & editing** — rope-backed text buffer, edit buffer, cursor/selection, editor view, clipboard.
4. **Rendering features** — scroll boxes, images (kitty/sixel), truecolor, double buffering diff, synchronized output.
5. **Input & terminal** — full key/mouse protocol, bracketed paste, focus management, terminal capability detection.
6. **Advanced** — animation timeline, syntax highlighting, 3D (optional), audio, SSH transport.
7. **Consumer** — port OpenCode TUI screens (session view, sidebar, dialogs, composer) onto gotui.

## OpenCode TUI mock (phase 1)
Implemented in `cmd/opencode-mock/main.go`:
- Header with title + spinner
- Sidebar with session list (focusable, ↑/↓, Enter)
- Chat area with message log + input
- Footer status bar
- Tab cycles focus; Ctrl+C quits

## Build
```
cd gotui
GOTOOLCHAIN=auto go build ./...
GOTOOLCHAIN=auto go run ./cmd/opencode-mock   # needs a TTY
```
