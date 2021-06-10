package panels

import "fyne.io/fyne/v2"

type Panel struct {
	Title, Intro string
	View func(w fyne.Window) fyne.CanvasObject
}

var (
	// Panels defines metada for each panel
	Panels = map[string]Panel{
		"home": {"Home", "", homeScreen},
		"configuration":{"Configuration",
			"Set up your Shortcast",
			configScreen},
	}

	// PanelIndex defines how the panels are laid out in the index tree
	PanelIndex = map[string][]string{
		"":            {"home", "configuration"},
	}
)