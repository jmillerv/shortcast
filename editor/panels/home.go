package panels

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func homeScreen(_ fyne.Window) fyne.CanvasObject {
	newFileButton := widget.NewButtonWithIcon("new", theme.DocumentCreateIcon(), func(){
		fmt.Println("creating new file")
	})
	openFileButton := widget.NewButtonWithIcon("open", theme.FileTextIcon(), func(){
		fmt.Println("opening file")
	})
	saveButton := widget.NewButtonWithIcon("save", theme.DocumentSaveIcon(), func(){
		fmt.Println("saving file")
	})

	bottomButtons := container.NewGridWithColumns(3, newFileButton, openFileButton, saveButton)
	return container.NewBorder(createToolBar(), bottomButtons, nil, nil)
}

func createToolBar() fyne.CanvasObject {
	t := widget.NewToolbar(widget.NewToolbarAction(BoldIcon(), func() { fmt.Println("Bold") }),
		widget.NewToolbarAction(ItalicIcon(), func() { fmt.Println("italic") }),
		widget.NewToolbarAction(StrikethroughIcon(), func() { fmt.Println("strikethrough") }),
		widget.NewToolbarAction(LinkIcon(), func() { fmt.Println("link") }),
		widget.NewToolbarAction(TitleIcon(), func() { fmt.Println("header") }),
		widget.NewToolbarAction(theme.FileImageIcon(), func() { fmt.Println("image") }),
		widget.NewToolbarAction(ListIcon(), func() { fmt.Println("list") }),
		widget.NewToolbarAction(NumberedListIcon(), func() { fmt.Println("numbered list") }),
		widget.NewToolbarAction(theme.CheckButtonIcon(), func() { fmt.Println("tasklist") }),
	)
	return t
}
