package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/radical-egg/zeepo/internal/tui"
	"github.com/rivo/tview"
)

func main() {
	grid := tui.GenerateColorGrid()
	app := tview.NewApplication()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			grid = tui.GenerateColorGrid()
			if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
				panic(err)
			}
		case tcell.KeyCtrlC:
			app.Stop()
		}
		return event
	})

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}