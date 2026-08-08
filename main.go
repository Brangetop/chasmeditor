package main

import (
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	editor := tview.NewApplication()
	editorArea := tview.NewTextArea()

	rootDir := "/home"
	root := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)
	explorer := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	add := func(target *tview.TreeNode, path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(filepath.Join(path, file.Name())).
				SetSelectable(file.IsDir())
			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node)
		}
	}

	add(root, rootDir)

	explorer.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			path := reference.(string)
			add(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})

	explorer.SetBorder(true)
	explorer.SetTitle("Explorer")

	editorArea.SetBorder(true)
	editorArea.SetTitle("Editor")

	flex := tview.NewFlex().
		AddItem(explorer, 0, 1, true).
		AddItem(editorArea, 0, 4, false)

	// logic of switching focus from one window to another
	elements := []tview.Primitive{explorer, editorArea}
	activeWindow := 0

	editor.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			activeWindow = (activeWindow + 1) % len(elements)
			editor.SetFocus(elements[activeWindow])
			return nil
		}

		return event
	})

	//starting the application
	if err := editor.SetRoot(flex, true).SetFocus(explorer).Run(); err != nil {
		panic(err)
	}

}
