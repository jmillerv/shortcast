package panels

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/shortcast/backend"
	"github.com/jmillerv/shortcast/backend/hugo"
)




func manageScreen(_ fyne.Window) fyne.CanvasObject {
	// CAST
	startCastButton := widget.NewButtonWithIcon("start cast", theme.MediaPlayIcon(), backend.StartShortcast)
	stopCastButton := widget.NewButtonWithIcon("stop cast", theme.MediaStopIcon(), backend.StopShortcast)
	castControls := container.NewGridWithColumns(2, stopCastButton, startCastButton)

	// POST
	newPostButton := container.NewVBox(
		container.NewHBox( // This cannot be the best practice way to get tiny buttons
		widget.NewButtonWithIcon("New", theme.ContentAddIcon(), hugo.AddPost)))
	postSidePanel := container.NewGridWithRows(1, newPostButton)
	postList := widget.NewList(getPostCount, createPostRow, updatePostRow)
	postTabContent := container.NewBorder(nil, nil, postSidePanel,postList)

	// THEME
	newThemeButton := container.NewVBox(
		container.NewHBox(
		widget.NewButtonWithIcon("Upload", theme.UploadIcon(), hugo.AddTheme)))
	themeSidePanel := container.NewGridWithRows(1, newThemeButton)
	themeList := widget.NewList(getThemeCount, createThemeRow, updateThemeRow)
	themeTabContent := container.NewBorder(nil, nil, themeSidePanel, themeList)

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Posts", theme.DocumentIcon(), postTabContent),
		container.NewTabItemWithIcon("Themes", theme.ColorPaletteIcon(), themeTabContent),
	)
	return container.NewBorder(nil, castControls, tabs, nil)
}

// getPostCount calls the backend API to retrieve the posts and then returns the number of posts. Inside of the function
// we determine the number of rows based on posts returned but the number of columns is determined by the current UI
// implementation
func getPostCount()int {
	return 5
}

// createPostRow generates a row with the post title and
func createPostRow() fyne.CanvasObject {
	icon := widget.NewIcon(theme.FileIcon())
	postTitle := widget.NewLabel("Test Post")
	editButton := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func(){
		fmt.Println("editing file")
	})
	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func(){
		fmt.Println("deleting file")
	})
	return container.NewHBox(icon, postTitle, layout.NewSpacer(), editButton, deleteButton)
}

func updatePostRow(id widget.ListItemID, template fyne.CanvasObject) {
	//label := template.(*widget.Label)
	//label.SetText("Update Post")
}

func getThemeCount()int {
	return 2
}

// createPostRow generates a row with the post title and
func createThemeRow() fyne.CanvasObject {
	postTitle := widget.NewLabel("Test Theme")
	selectButton := widget.NewButtonWithIcon("", theme.VisibilityIcon(), func(){
		fmt.Println("setting theme")
	})
	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func(){
		fmt.Println("deleting theme")
	})
	return container.NewHBox(postTitle, layout.NewSpacer(), selectButton, deleteButton)
}

func updateThemeRow(id widget.ListItemID, template fyne.CanvasObject) {
	//label := template.(*widget.Label)
	//label.SetText("Update Post")
}

