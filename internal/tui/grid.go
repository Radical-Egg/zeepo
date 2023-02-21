package tui

import (
	"github.com/radical-egg/zeepo/internal/services"
	"github.com/rivo/tview"
)

func GenerateColorGrid() *tview.Grid {
	colors := services.GetColorPallete()

	prims := make(map[string]tview.Primitive, len(colors.Result))

	for count, color := range colors.Result {
		hex := services.ConvRGBToHex(color...)
		prims["prim" + string(count)] = newPrimitive(hex, color...)
	}

	grid := tview.NewGrid().
		SetRows(0, 0, 0).
		SetColumns(0, 0, 0, 0, 0)
		
	grid.SetBorder(true)
	
	counter := 0

	for _, prim := range prims {
		grid.AddItem(prim, 0, counter, 5, 1, 0, 0, false)
		counter++
	}

	return grid
}