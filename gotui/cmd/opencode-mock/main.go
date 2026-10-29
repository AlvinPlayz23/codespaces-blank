package main

import (
	"fmt"

	"gotui/component"
	"gotui/style"

	"gotui"
)

func main() {
	root := component.NewBox(component.Column)
	root.Bg = style.BgRGB(18, 18, 24)
	root.SetPrefH(0)
	root.SetFlex(1)

	// header
	header := component.NewBox(component.Row)
	header.SetPrefH(3)
	header.Border = true
	header.BorderColor = style.Cyan
	header.Bg = style.BgRGB(24, 24, 32)
	header.Title = "opencode"
	headerText := component.NewText("opencode  •  gotui mock (Go port)", style.Style{Fg: style.BrightCyan, Bold: true})
	headerText.SetFlex(1)
	header.Add(headerText)
	sp := component.NewSpinner("ready")
	sp.SetPrefW(12)
	header.Add(sp)
	root.Add(header)

	// body
	body := component.NewBox(component.Row)
	body.SetFlex(1)
	root.Add(body)

	// sidebar
	sidebar := component.NewBox(component.Column)
	sidebar.SetPrefW(28)
	sidebar.Border = true
	sidebar.BorderColor = style.Gray
	sidebar.Title = "Sessions"
	sessions := []component.ListItem{
		{Label: "main", Value: "main"},
		{Label: "refactor-tui", Value: "refactor"},
		{Label: "bugfix-42", Value: "bugfix"},
		{Label: "docs", Value: "docs"},
	}
	sessList := component.NewList(sessions)
	sessList.SetFlex(1)
	sidebar.Add(sessList)
	body.Add(sidebar)

	// main chat
	mainBox := component.NewBox(component.Column)
	mainBox.SetFlex(1)
	mainBox.Border = true
	mainBox.BorderColor = style.Gray
	mainBox.Title = "Chat"
	body.Add(mainBox)

	chat := component.NewBox(component.Column)
	chat.SetFlex(1)
	chat.PadX = 1
	chat.PadY = 1
	mainBox.Add(chat)

	// initial messages
	addMsg(chat, "assistant", "Hello! This is a mock OpenCode TUI running on gotui (Go).")
	addMsg(chat, "assistant", "Type a message and press Enter. Tab switches focus.")

	input := component.NewInput()
	input.Placeholder = "Ask anything..."
	input.SetPrefH(1)
	mainBox.Add(input)

	input.OnSubmit = func(v string) {
		if v == "" {
			return
		}
		addMsg(chat, "user", v)
		input.Value = ""
		input.Cursor = 0
		addMsg(chat, "assistant", "Echo: "+v)
	}

	// footer
	footer := component.NewBox(component.Row)
	footer.SetPrefH(1)
	footer.Bg = style.BgRGB(24, 24, 32)
	status := component.NewText(" Tab: focus  ↑↓: navigate  Enter: send  Ctrl+C: quit", style.Style{Fg: style.Gray})
	status.SetFlex(1)
	footer.Add(status)
	root.Add(footer)

	app := gotui.NewApp(root)
	app.Focused = input
	input.Focused = true

	if err := app.Run(); err != nil {
		fmt.Println("error:", err)
	}
}

func addMsg(chat *component.Box, role, text string) {
	color := style.White
	prefix := "assistant: "
	if role == "user" {
		color = style.BrightGreen
		prefix = "you: "
	} else {
		color = style.BrightCyan
	}
	t := component.NewText(prefix+text, style.Style{Fg: color})
	t.SetPrefH(1 + (len(text)+20)/40) // rough
	chat.Add(t)
}
