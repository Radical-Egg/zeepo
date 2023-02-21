package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/radical-egg/zeepo/internal/tui"
	"github.com/rivo/tview"
)

func main() {
	pages := tview.NewPages()
	grid := tui.GenerateColorGrid()
	app := tview.NewApplication()

	pages.AddPage("Color Grid", grid, true, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
			case tcell.KeyEnter:
				grid = tui.GenerateColorGrid()
				pages.RemovePage("Color Grid")
				pages.AddPage("Color Grid", grid, true, true)
				pages.SwitchToPage("Color Grid")
			case tcell.KeyCtrlC, tcell.KeyEsc:
				app.Stop()
			}
		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}