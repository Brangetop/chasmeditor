package main

import (
	"github.com/gdamore/tcell/v2"
)

func main() {
	app := &App{}
	app.Init()

	app.tviewApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// temporary debugging stuff
		/*app.ChangeStatus(
			fmt.Sprintf("key=%v rune=%q mods=%v", event.Key(), event.Rune(), event.Modifiers()),
		) */
		/*
			modCtrl := event.Modifiers()&tcell.ModCtrl != 0
			modShift := event.Modifiers()&tcell.ModShift != 0
			r := event.Rune()

			if modCtrl && modShift && (r == 's' || r == 'S') {
				_ = app.SaveFileAs(app.currentPath)
				return nil
			}

				// DO 3 key modifiers even work???

		*/

		switch event.Key() {
		case tcell.KeyCtrlO:
			app.activeWindow = (app.activeWindow + 1) % len(app.elements)
			app.tviewApp.SetFocus(app.elements[app.activeWindow])
			return nil

		// Exiting
		case tcell.KeyCtrlQ:
			//app.ExitWithSaving(app.currentPath)

			app.tviewApp.Stop()
			return nil

		// Saving and write like in nano, add next day
		case tcell.KeyCtrlS:
			if err := app.SaveFile(app.currentPath); err != nil {
				app.ChangeStatus(err.Error())
			}
			return nil

		// Copying
		case tcell.KeyCtrlC:
			return nil

		case tcell.KeyCtrlB:
			app.ShowDirectoryTree(!app.direcorryTreeVisible)
			return nil
		}

		return event
	})

	//starting the application
	if err := app.tviewApp.SetRoot(app.root, true).SetFocus(app.explorer).Run(); err != nil {
		app.ChangeStatus(err.Error())
	}
}
