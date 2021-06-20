package utils

import "unicode"

func AddSpacesToPascalString(s string) string {
	var spacedString string
	for i := 0; i < len(s); i++ {
		if i != 0 && unicode.IsUpper(rune(s[i])){
			spacedString = spacedString + " " + string(s[i])
			continue
		}
		spacedString = spacedString + string(s[i])
	}
	return spacedString
}
