package panels

import "fyne.io/fyne/v2"

type Panel struct {
	Title, Intro string
	View func(w fyne.Window) fyne.CanvasObject
}

var (
	// Panels defines metada for each panel
	Panels = map[string]Panel{
		"home": {"Shortcast", "", homeScreen},
		"getting started": {"Getting Started", "Help for a successful experience", gettigStartedScreen},
		"admin": {"Admin", "Manage your Shortcast", adminScreen},
		"configuration":{"Configuration",
			"Set up your Shortcast",
			configScreen},
		"documentation": {"Documentation", "Guides and Best Practices", documentationScreen},

	}

	// PanelIndex defines how the panels are laid out in the index tree
	PanelIndex = map[string][]string{
		"":            {"home", "getting started", "admin", "configuration", "documentation"},
	}
)