package panels

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

var DisableTutorialSeen = binding.NewBool()

func gettigStartedScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewHBox()
}

func SetGettingStartedBool(update bool) {
	err := DisableTutorialSeen.Set(update)
	if err != nil {
		fmt.Println("error setting boolean")
	}
	currentSetting, _ :=  DisableTutorialSeen.Get()
	fmt.Println("getting started screen ",currentSetting)
}