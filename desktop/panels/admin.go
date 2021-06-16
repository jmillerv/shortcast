package panels

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func adminScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewHBox()
}


// admin will have the ability to manage the website name and preferred themes