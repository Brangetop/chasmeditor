package main

import (
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	editor := tview.NewApplication()

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

	editorBox := tview.NewBox().
		SetBorder(true).
		SetTitle("Editor")

	flex := tview.NewFlex().
		AddItem(explorer, 0, 1, true).
		AddItem(editorBox, 0, 4, false)

	if err := editor.SetRoot(flex, true).SetFocus(explorer).Run(); err != nil {
		panic(err)
	}

}
