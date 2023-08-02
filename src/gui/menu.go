package gui

import (
	"os"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialogue"
)

func menu() {
	//This is the main GUI window. Every GUI component will be linked through here.
	menu := app.New()
	menu_window := menu.NewWindow("Container")

}