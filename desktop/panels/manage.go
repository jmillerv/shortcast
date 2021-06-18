package panels

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/shortcast/backend"
	"github.com/jmillerv/shortcast/backend/posts"
)




func manageScreen(_ fyne.Window) fyne.CanvasObject {
	startCastButton := widget.NewButtonWithIcon("start cast", theme.MediaPlayIcon(), backend.StartShortcast)
	stopCastButton := widget.NewButtonWithIcon("stop cast", theme.MediaStopIcon(), backend.StopShortcast)
	castControls := container.NewGridWithColumns(2, stopCastButton, startCastButton)
	newPostButton := widget.NewButtonWithIcon("New", theme.ContentAddIcon(), posts.Add)
	postSidePanel := container.NewGridWithRows(1, newPostButton)
	newThemeButton := widget.NewButtonWithIcon("Upload", theme.UploadIcon(), posts.Add)
	themeSidePanel := container.NewGridWithRows(1, newThemeButton)

	postTabContent := container.New(layout.NewFormLayout(), postSidePanel, widget.NewLabel("list of posts with edit/remove buttons here"))
	themeTabContent := container.New(layout.NewFormLayout(), themeSidePanel, widget.NewLabel("list of themes with edit remove"))

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Posts", theme.DocumentIcon(), postTabContent),
		container.NewTabItemWithIcon("Themes", theme.ColorPaletteIcon(), themeTabContent),
	)
	return container.NewBorder(nil, castControls, tabs, nil)
}


// admin will have the ability to manage the website name and preferred themes