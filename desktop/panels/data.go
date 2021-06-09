package content

import "fyne.io/fyne/v2"

type Panel struct {
	Title, Intro string
	View func(w fyne.Window) fyne.CanvasObject
}

var (
	// Panels defines metada for each panel
	Panels = map[string]Panel{

	}
)