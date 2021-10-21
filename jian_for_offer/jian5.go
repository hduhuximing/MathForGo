package jian_for_offer

import "strings"

func replaceSpace(s string) string {
	var stringBuilder strings.Builder
	for i, element := range s {
		if element == ' ' {
			stringBuilder.WriteString("%20")
		} else {
			stringBuilder.WriteByte(s[i])
		}
	}
	return stringBuilder.String()
}
