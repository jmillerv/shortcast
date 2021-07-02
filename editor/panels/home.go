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
		widget.NewToolbarAction(theme.InfoIcon(), func(){ toolBarInfoScreen()}),
	)
	return t
}

func toolBarInfoScreen() {
	row := make([]fyne.CanvasObject, len(iconConsts))
	w := fyne.CurrentApp().NewWindow("Toolbar Info")
	for i := range iconConsts {
		title := widget.NewLabel(IconInfoMap[iconConsts[i]])
		iconFunc := IconFuncMap[iconConsts[i]]
		icon := iconFunc()
		row[i] = container.NewHBox(widget.NewIcon(icon), layout.NewSpacer(), title)
	}
	list := widget.NewList(getToolbarInfoLength, func() fyne.CanvasObject {
		return container.NewGridWithColumns(1, row...)
	}, updateToolBarRow)
	w.Resize(fyne.NewSize(240, 480))
	w.SetContent(list)
	w.Show()
}

func getToolbarInfoLength() int {
	return 1
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
