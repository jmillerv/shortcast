package post

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"runtime"
)

// for reference https://github.com/fynelabs/notes/blob/main/note.go

type Post struct {
	Content binding.String
	Entry *widget.Entry
}

func New() *Post {
	p := new(Post)
	err := p.Content.Set(placeHolderContent())
	if err != nil {
		return nil
	}
	return new(Post)
}

func (p *Post) Save() {

}

func (p *Post) Edit() {
	p.Entry.Bind(p.Content)
	p.Entry.Validator = nil
}

func (p *Post) Delete() {

}

func placeHolderContent() string {
	text := "Welcome! type to add text."
	if fyne.CurrentDevice().HasKeyboard() {
		modifier := "ctrl"
		if runtime.GOOS == "darwin" {
			modifier = "cmd"
		}
		text += fmt.Sprintf("\n\nOr use the keyboard shortcut %s+N.", modifier)
	}
	return text
}