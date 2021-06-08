package main

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	window := a.NewWindow("Shortcast")
	tabs := container.NewAppTabs()
	tabs.Append(container.NewTabItemWithIcon(
		"Dashboard", theme.HomeIcon(), widget.NewLabel("Check status and launch"),
		))
	tabs.Append(container.NewTabItemWithIcon(
		"Configuration", theme.SettingsIcon(), widget.NewLabel("settings"),
	))
	tabs.SetTabLocation(container.TabLocationLeading)
	window.SetContent(tabs)
	window.Resize(fyne.Size{
		Width:  600,
		Height: 400,
	})
	window.ShowAndRun()
}

