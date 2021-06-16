package panels

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

var gettingStartedSeen = binding.NewBool()

func gettigStartedScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewHBox()
}