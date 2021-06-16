package panels

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testHomeDirectory      = "HomeDirectory"
	testShortcastDirectory = "ShortcastDirectory"
	testWifiName           = "WifiName"
	testPassword           = "Password"
	testPasswordRequired   = "PasswordRequired"
)

func TestSortStruct(t *testing.T) {
	formData := binding.BindStruct(&configStruct{})
	keys := sortStruct(formData.Keys())
	assert.Equal(t, keys[0], testHomeDirectory)
	assert.Equal(t, keys[1], testShortcastDirectory)
	assert.Equal(t, keys[2], testWifiName)
	assert.Equal(t, keys[3], testPassword)
	assert.Equal(t, keys[4], testPasswordRequired)
}
