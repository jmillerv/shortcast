package panels

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func homeScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to Shortcast", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewHBox(
			// TODO hyper links won't work because it's not intended to be connected to internet
			widget.NewHyperlink("about",  parseURL("https://github.com/jmillerv")),
			widget.NewLabel("-"),
			widget.NewHyperlink("report issue", parseURL("https://github.com/jmillerv")),
			widget.NewLabel("-"),
			widget.NewHyperlink("contribute", parseURL("https:/github.com/jmillerv")),
		),
		))
}