package main

import (
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	tviewApp   *tview.Application
	explorer   *tview.TreeView
	editorArea *tview.TextArea
	statusBar  *tview.TextView

	activeWindow int
	elements     []tview.Primitive

	status string
}

func (a *App) Init() {
	rootDir := "."

	a.tviewApp = tview.NewApplication()
	a.explorer = tview.NewTreeView()
	a.editorArea = tview.NewTextArea()
	a.statusBar = tview.NewTextView()

	a.activeWindow = 0
	a.elements = []tview.Primitive{a.explorer, a.editorArea}
	a.status = "Initialization complete"

	root := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)
	a.explorer.SetRoot(root).SetCurrentNode(root)

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
	// root.Expand()

	a.explorer.SetSelectedFunc(func(node *tview.TreeNode) {
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

	a.explorer.SetBorder(true)
	a.explorer.SetTitle("Explorer")

	a.editorArea.SetBorder(true)
	a.editorArea.SetTitle("Editor")

	a.statusBar.SetDynamicColors(true)
	a.statusBar.SetRegions(true)
	a.statusBar.SetTextAlign(tview.AlignLeft)
	a.statusBar.SetBorder(true)
}
