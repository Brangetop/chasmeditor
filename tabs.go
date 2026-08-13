package main

import "github.com/rivo/tview"

type Tab struct {
	path  string
	title string
	text  string
	dirty bool
}

type Tabs struct {
	app *App

	tabBar *tview.List
	editor *tview.TextArea
	flex   *tview.Flex

	tabs      []Tab
	activeIdx int
}

func (t *Tabs) NewTab() {

}
