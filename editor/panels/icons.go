package panels

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const (
	IconNameBold          fyne.ThemeIconName = "bold"
	IconNameCode          fyne.ThemeIconName = "code"
	IconNameImage         fyne.ThemeIconName = "image"
	IconNameItalic        fyne.ThemeIconName = "italic"
	IconNameLink          fyne.ThemeIconName = "link"
	IconNameList          fyne.ThemeIconName = "list"
	IconNameNumberedList  fyne.ThemeIconName = "numbered_list"
	IconNameStrikethrough fyne.ThemeIconName = "strikethrough"
	IconNameTasklist      fyne.ThemeIconName = "tasklist"
	IconNameTableChart    fyne.ThemeIconName = "table_chart"
	IconNameTitle         fyne.ThemeIconName = "title"
)

var iconConsts = []fyne.ThemeIconName{
	IconNameBold,
	IconNameCode,
	IconNameImage,
	IconNameItalic,
	IconNameLink,
	IconNameList,
	IconNameNumberedList,
	IconNameStrikethrough,
	IconNameTasklist,
	IconNameTableChart,
	IconNameTitle,
}

var icons = map[fyne.ThemeIconName]fyne.Resource{
	IconNameBold:          theme.NewThemedResource(formatBoldRes),
	IconNameCode:          theme.NewThemedResource(codeRes),
	IconNameImage:         theme.NewThemedResource(imageIconRes),
	IconNameItalic:        theme.NewThemedResource(formatItalicRes),
	IconNameLink:          theme.NewThemedResource(addLinkRes),
	IconNameList:          theme.NewThemedResource(listRes),
	IconNameNumberedList:  theme.NewThemedResource(formatNumberedListRes),
	IconNameStrikethrough: theme.NewThemedResource(strikethroughRes),
	IconNameTasklist:      theme.NewThemedResource(taskListRes),
	IconNameTableChart:    theme.NewThemedResource(tableChartRes),
	IconNameTitle:         theme.NewThemedResource(titleRes),
}

var IconFuncMap = map[fyne.ThemeIconName]func() fyne.Resource{
	IconNameBold:          BoldIcon,
	IconNameCode:          CodeIcon,
	IconNameImage:         ImageIcon,
	IconNameItalic:        ItalicIcon,
	IconNameLink:          LinkIcon,
	IconNameList:          ListIcon,
	IconNameNumberedList:  NumberedListIcon,
	IconNameStrikethrough: StrikethroughIcon,
	IconNameTasklist:      TaskListIcon,
	IconNameTableChart:    TableChartIcon,
	IconNameTitle:         TitleIcon,
}

var IconInfoMap = map[fyne.ThemeIconName]string{
	IconNameBold:          "Bold",
	IconNameCode:          "Code",
	IconNameImage:         "Add Image",
	IconNameItalic:        "Italic",
	IconNameLink:          "Link",
	IconNameList:          "List",
	IconNameNumberedList:  "Numbered list",
	IconNameStrikethrough: "Strikethrough",
	IconNameTasklist:      "Task list",
	IconNameTableChart:    "Table",
	IconNameTitle:         "Heading",
}

// may need to implement a safeIconLookup function like one that exists in fyne/theme/icons.go

func BoldIcon() fyne.Resource {
	return icons[IconNameBold]
}

func CodeIcon() fyne.Resource {
	return icons[IconNameCode]
}

func ImageIcon() fyne.Resource {
	return icons[IconNameImage]
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

func TaskListIcon() fyne.Resource {
	return icons[IconNameTasklist]
}

func TableChartIcon() fyne.Resource {
	return icons[IconNameTableChart]
}

func TitleIcon() fyne.Resource {
	return icons[IconNameTitle]
}
