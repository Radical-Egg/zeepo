package tui

import (
	"github.com/rivo/tview"
)

func HelpMenu() *tview.TextView {
	t := tview.NewTextView().
	SetDynamicColors(true).
	SetTextAlign(tview.AlignCenter).
	SetText(`
		<Enter> Refresh Color Palette
		<ESC> Close the program
		<TAB> Show this menu
	`)

	t.SetBorder(true)
	t.SetTitle("Help Menu")
	t.SetTextAlign(tview.AlignCenter)

	return t
}