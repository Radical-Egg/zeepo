package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func newPrimitive(text string, colors ...int) tview.Primitive {
	color := tcell.NewRGBColor(int32(colors[0]), int32(colors[1]), int32(colors[2]))

	t := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText(text)
	t.SetBackgroundColor(color)
	

	return t;
}