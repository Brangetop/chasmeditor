package main

import (
	"github.com/gdamore/tcell/v2"
)

func main() {
	app := &App{}
	app.Init()

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

	//starting the application
	if err := app.tviewApp.SetRoot(app.root, true).SetFocus(app.explorer).Run(); err != nil {
		app.ChangeStatus(err.Error())
	}
}
