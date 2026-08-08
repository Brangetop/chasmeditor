package main

import (
	"github.com/rivo/tview"
)

func main() {
	editor := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Explorer"), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Editor"), 0, 4, false)

	if err := editor.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}

}
