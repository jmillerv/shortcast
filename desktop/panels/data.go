package panels

import "fyne.io/fyne/v2"

type Panel struct {
	Title string
	View func(w fyne.Window) fyne.CanvasObject
}

var (
	// Panels defines metadata for each panel
	Panels = map[string]Panel{
		"home": {"Shortcast", homeScreen},
		"getting started": {"Getting Started", gettigStartedScreen},
		"configuration":{"Configuration", configScreen},
		"documentation": {"Documentation",  documentationScreen},
		"manage": {"Manage", manageScreen},
	}

	// PanelIndex defines how the panels are laid out in the index tree
	PanelIndex = map[string][]string{
		"":            {"home", "getting started", "configuration", "documentation", "manage"},
	}
)