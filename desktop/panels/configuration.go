package panels

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"os"
)

const (
	defaultHomeDir          = "/home/pi"
	defaultShortcastDir     = "/shortcast"
	defaultThemesDir        = "/themes"
	defaultWifiName         = "Shortcast"
	defaultPassword         = "cageceramics"
	defaultPasswordRequired = false

	// environment variables
	shortcastHomeDir   = "SHORTCAST_HOME_DIR"
	shortcastWorkDir   = "SHORTCAST_WORK_DIR"
	shortcastThemesDir = "SHORTCAST_THEMES_DIR"
	shortcastWifiName  = "SHORTCAST_WIFI_NAME"
)

type configStruct struct {
	HomeDirectory      string `json:"home_directory"`
	ShortcastDirectory string `json:"shortcast_directory"`
	WifiName           string `json:"wifi_name"`
	Password           string `json:"password"`
	PasswordRequired   bool   `json:"password_required"`
}

func configScreen(win fyne.Window) fyne.CanvasObject {
	formStruct := configStruct{
		HomeDirectory:      defaultHomeDir,
		ShortcastDirectory: defaultShortcastDir,
		WifiName:           defaultWifiName,
		Password:           defaultPassword,
		PasswordRequired:   defaultPasswordRequired,
	}

	formData := binding.BindStruct(&formStruct)
	form := newFormWithData(formData)
	form.OnSubmit = func() {
		// createSettings()
		fmt.Println("Struct:\n", formData)
	}
	infoButtons := getInfoButtons(win)
	return container.NewBorder(nil, nil, nil, nil, container.NewHSplit(form, infoButtons))

}

func getInfoButtons(win fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		container.NewGridWithRows(5,
			widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
				dialog.ShowInformation("Home Directory", "This is your device's home folder", win)
			}),
			widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
				dialog.ShowInformation("Shortcast Directory", "This is where shortcast related files will live on the pi", win)
			}),
			widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
				dialog.ShowInformation("WiFi Name", "This is the connection name that will be visible to users", win)
			}),
			widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
				dialog.ShowInformation("Password", "This is where you set a password", win)
			}),
			widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
				dialog.ShowInformation("Password Required", "Click if you want to enable who can access your Shortcast", win)
			}),
			),
		)
}

func newFormWithData(data binding.DataMap) *widget.Form {
	keys := sortStruct(data.Keys())
	items := make([]*widget.FormItem, len(keys))
	for i, k := range keys {
		data, err := data.GetItem(k)
		if err != nil {
			items[i] = widget.NewFormItem(k, widget.NewLabel(err.Error()))
		}
		items[i] = widget.NewFormItem(k, createBoundItem(data))
	}
	return widget.NewForm(items...)
}

func createBoundItem(v binding.DataItem) fyne.CanvasObject {
	switch val := v.(type) {
	case binding.Bool:
		return widget.NewCheckWithData("", val)
	case binding.Float:
		s := widget.NewSliderWithData(0, 1, val)
		s.Step = 0.01
		return s
	case binding.Int:
		return widget.NewEntryWithData(binding.IntToString(val))
	case binding.String:
		return widget.NewEntryWithData(val)
	default:
		return widget.NewLabel("")
	}
}

// sortStruct orders the configStruct to ensure a consistent layout
func sortStruct(x []string) []string {
	keys := make([]string, len(x))
	for _, k := range x {
		switch k {
		case "HomeDirectory":
			keys[0] = k
		case "ShortcastDirectory":
			keys[1] = k
		case "WifiName":
			keys[2] = k
		case "Password":
			keys[3] = k
		case "PasswordRequired":
			keys[4] = k
		default:
			keys = append(keys, k)
		}
	}
	return keys
}

// createSettings sets wifi settings as environment variables
func createSettings(data binding.Struct) {
	//home, err := data.GetItem("HomeDirectory")
	//if err != nil {
	//	log.Println(err)
	//}
	err := os.Setenv("SHORTCAST_HOME_DIR", "")
	if err != nil {
		return
	}
}
