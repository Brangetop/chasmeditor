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
		case tcell.KeyEsc:
			app.activeWindow = (app.activeWindow + 1) % len(app.elements)
			app.tviewApp.SetFocus(app.elements[app.activeWindow])
			return nil

		// Exiting
		case tcell.KeyCtrlQ:
			app.tviewApp.Stop()
			return nil

		// Saving
		case tcell.KeyCtrlS:
			if err := app.SaveFile(app.currentPath); err != nil {
				app.ChangeStatus(err.Error())
			}
			return nil

		// Copying
		case tcell.KeyCtrlC:
			// need to find a way to work with system-wide clipboard, not managing the local one
			// app.CopySelection()
			return nil
		}

		return event
	})

	root := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(flex, 0, 1, true).
		AddItem(app.statusBar, 4, 0, false)

	//starting the application
	if err := app.tviewApp.SetRoot(root, true).SetFocus(app.explorer).Run(); err != nil {
		app.ChangeStatus(err.Error())
	}
}
