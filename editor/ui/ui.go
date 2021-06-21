package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/shortcast/editor/panels"
)

const (
	preferenceCurrentPanel = "currentPanel"
)


var topWindow fyne.Window

func Render() {
	a := app.NewWithID("editor.app")
	w := a.NewWindow("Markdown Editor")

	topWindow = w

	w.SetMaster()

	content := container.NewMax()

	title := widget.NewLabel("Component name")
	setPanel := func(p panels.Panel) {
		if fyne.CurrentDevice().IsMobile() {
			child := a.NewWindow(p.Title)
			topWindow = child
			child.SetContent(p.View(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}
		title.SetText(p.Title)

		content.Objects = []fyne.CanvasObject{p.View(w)}
		content.Refresh()
	}

	panel := container.NewBorder(container.NewVBox(title, widget.NewSeparator()), nil, nil, nil, content)
	if fyne.CurrentDevice().IsMobile() {
		w.SetContent(createNav(setPanel, false))
	} else {
		split := container.NewHSplit(createNav(setPanel, true), panel)
		split.Offset = 0.2
		w.SetContent(split)
	}

	w.Resize(fyne.Size{Width: 640, Height: 460})
	w.ShowAndRun()
}

func createNav(setPanel func(panel panels.Panel), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return panels.PanelIndex[uid]
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		IsBranch: func(uid string) bool {
			children, ok := panels.PanelIndex[uid]
			return ok && len(children) > 0
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			p, ok := panels.Panels[uid]
			if !ok {
				fyne.LogError(fmt.Sprintf("Missing panel %s", uid), nil)
			}
			obj.(*widget.Label).SetText(p.Title)
		},
		OnSelected: func(uid string) {
			if p, ok := panels.Panels[uid]; ok {
				a.Preferences().SetString(preferenceCurrentPanel, uid)
				setPanel(p)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentPanel, "home")
		tree.Select(currentPref)
	}

	return container.NewBorder(nil, nil, nil, nil, tree)
}