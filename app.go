package main

import (
	"errors"
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

	status      string
	currentPath string

	// clipboard string
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
			a.ChangeStatus(err.Error())
		}
		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(filepath.Join(path, file.Name())).
				// only directories are selecrable
				//SetSelectable(file.IsDir())

				// everything is selectable
				SetSelectable(true)
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
		path := reference.(string)
		info, err := os.Stat(path)
		if err != nil {
			a.ChangeStatus(err.Error())
			return
		}

		if info.IsDir() {
			// lazy collapse
			// redo in the nearest future with directory refreshing
			children := node.GetChildren()
			if len(children) == 0 {
				// Load and show files in this directory.
				path := reference.(string)
				add(node, path)
			} else {
				// Collapse if visible, expand if collapsed.
				node.SetExpanded(!node.IsExpanded())
			}

			return
		}

		// if not a dir => is a file
		/* text, err := loadFileText(path)
		if err != nil {
			a.ChangeStatus(err.Error())
			return
		}
		*/

		// loading file
		a.LoadFile(path)
	})

	a.explorer.SetBorder(true)
	a.explorer.SetTitle("Explorer")

	a.editorArea.SetBorder(true)
	a.editorArea.SetTitle("Editor")

	a.statusBar.SetDynamicColors(true)
	a.statusBar.SetRegions(true)
	a.statusBar.SetTextAlign(tview.AlignLeft)
	a.statusBar.SetBorder(true)
	// a.statusBar.SetTitle("Status")

	a.ChangeStatus("Initialization complete")
}

func (a *App) SetPath(path string) {
	a.currentPath = path
}

func (a *App) ChangeStatus(st string) {
	a.status = st
	a.statusBar.SetText(a.status)
}

// maybe remove or implement similar method for saving
func (a *App) ChangeStatusOpenFile(path string) {
	a.status = "Opened file: " + path
	a.statusBar.SetText(a.status)
}

func (a *App) LoadFile(path string) {
	text, err := loadFileText(path)
	if err != nil {
		a.ChangeStatus(err.Error())
		return
	}
	a.editorArea.SetText(text, false)
	a.editorArea.SetTitle(filepath.Base(path))

	a.ChangeStatusOpenFile(path)
	a.SetPath(path)
}

func (a *App) SaveFile(path string) error {
	if path == "" {
		return errors.New("Not a valid path")
	}
	text := a.editorArea.GetText()

	a.ChangeStatus("Saved file: " + path)
	return os.WriteFile(path, []byte(text), 0644)
}
