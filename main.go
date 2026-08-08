package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := &App{}
	app.Init()

	flex := tview.NewFlex().
		AddItem(app.explorer, 0, 1, true).
		AddItem(app.editorArea, 0, 4, false)

	app.tviewApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			app.activeWindow = (app.activeWindow + 1) % len(app.elements)
			app.tviewApp.SetFocus(app.elements[app.activeWindow])
			return nil
		}

		return event
	})

	//starting the application
	if err := app.tviewApp.SetRoot(flex, true).SetFocus(app.explorer).Run(); err != nil {
		panic(err)
	}

}
