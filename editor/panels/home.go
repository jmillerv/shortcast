package panels

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/shortcast/editor/format"
)

func createToolBar() fyne.CanvasObject {
	t := widget.NewToolbar(
		widget.NewToolbarAction(TitleIcon(), format.Heading),
		widget.NewToolbarAction(BoldIcon(), format.Bold),
		widget.NewToolbarAction(ItalicIcon(), format.Italic),
		widget.NewToolbarAction(StrikethroughIcon(), format.Strikethrough),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(ListIcon(), format.List),
		widget.NewToolbarAction(NumberedListIcon(), format.NumberedList),
		widget.NewToolbarAction(TaskListIcon(), format.TaskList),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(LinkIcon(), format.Link),
		widget.NewToolbarAction(CodeIcon(), format.Code),
		widget.NewToolbarAction(ImageIcon(), format.AddImage),
		widget.NewToolbarAction(TableChartIcon(), format.Table),
		widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.InfoIcon(), func(){toolBarInfo()}),
	)
	return t
}

func toolBarInfo() fyne.Window {
	w := fyne.CurrentApp().NewWindow("Toolbar Info")
	list := widget.NewList(getToolbarInfoLength, createToolBarRow, updateToolBarRow)
	w.SetContent(list)
	fmt.Println("returning toolbar info")
	return w
}

func getToolbarInfoLength() int {
	return len(IconFuncMap)
}

func createToolBarRow() fyne.CanvasObject {
	for _, i := range iconConsts {
		title := widget.NewLabel(IconInfoMap[i])
		iconFunc := IconFuncMap[i]
		icon := iconFunc()
		return container.NewHBox(widget.NewIcon(icon), layout.NewSpacer(), title)
	}
	return widget.NewLabel("")
}

func updateToolBarRow(id widget.ListItemID, template fyne.CanvasObject) {
	//label := template.(*widget.Label)
	//label.SetText("Update Post")
}

func homeScreen(_ fyne.Window) fyne.CanvasObject {
	newFileButton := widget.NewButtonWithIcon("new", theme.DocumentCreateIcon(), func() {
		fmt.Println("creating new file")
	})
	openFileButton := widget.NewButtonWithIcon("open", theme.FileTextIcon(), func() {
		fmt.Println("opening file")
	})
	saveButton := widget.NewButtonWithIcon("save", theme.DocumentSaveIcon(), func() {
		fmt.Println("saving file")
	})

	bottomButtons := container.NewGridWithColumns(3, newFileButton, openFileButton, saveButton)
	return container.NewBorder(createToolBar(), bottomButtons, nil, nil)
}
