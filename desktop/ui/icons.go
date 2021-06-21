package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const (
	IconNameBold          fyne.ThemeIconName = "bold"
	IconNameCode          fyne.ThemeIconName = "code"
	IconNameItalic        fyne.ThemeIconName = "italic"
	IconNameLink          fyne.ThemeIconName = "link"
	IconNameList          fyne.ThemeIconName = "list"
	IconNameNumberedList  fyne.ThemeIconName = "numbered list"
	IconNameStrikethrough fyne.ThemeIconName = "strikethrough"
	IconNameTableChart    fyne.ThemeIconName = "table chart"
	IconNameTitle         fyne.ThemeIconName = "title"
)

var icons = map[fyne.ThemeIconName]fyne.Resource{
	IconNameBold:          theme.NewThemedResource(formatBoldRes),
	IconNameCode:          theme.NewThemedResource(codeRes),
	IconNameItalic:        theme.NewThemedResource(formatItalicRes),
	IconNameLink:          theme.NewThemedResource(addLinkRes),
	IconNameList:          theme.NewThemedResource(listRes),
	IconNameNumberedList:  theme.NewThemedResource(formatNumberedListRes),
	IconNameStrikethrough: theme.NewThemedResource(strikethroughRes),
	IconNameTableChart:    theme.NewThemedResource(tableChartRes),
	IconNameTitle:         theme.NewThemedResource(titleRes),
}

// may need to implement a safeIconLookup function like one that exists in fyne/theme/icons.go

func BoldIcon() fyne.Resource {
	return icons[IconNameBold]
}

func CodeIcon() fyne.Resource {
	return icons[IconNameCode]
}

func ItalicIcon() fyne.Resource {
	return icons[IconNameItalic]
}

func LinkIcon() fyne.Resource {
	return icons[IconNameLink]
}

func ListIcon() fyne.Resource {
	return icons[IconNameList]
}

func NumberedListIcon() fyne.Resource {
	return icons[IconNameNumberedList]
}

func StrikethroughIcon() fyne.Resource {
	return icons[IconNameStrikethrough]
}

func TableChartIcon() fyne.Resource {
	return icons[IconNameTableChart]
}

func TitleIcon() fyne.Resource {
	return icons[IconNameTitle]
}
