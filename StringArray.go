package coresync

import "strings"

func StringToArray(s string) []string {
	sb := strings.Builder{}
	arr := make([]string, 0, 8)
	esc := false
	for _, c := range s {
		if esc {
			sb.WriteRune(c)
			esc = false
			continue
		}
		if c == '\\' {
			esc = true
			continue
		}
		if c == '|' {
			arr = append(arr, sb.String())
			sb = strings.Builder{}
			continue
		}
		sb.WriteRune(c)
	}
	if sb.Len() > 0 {
		arr = append(arr, sb.String())
	}
	return arr
}

func ArrayToString(array []string) string {
	arr := make([]string, 0, 8)
	for _, s := range array {
		s = strings.ReplaceAll(s, "\\", "\\\\")
		s = strings.ReplaceAll(s, "|", "\\|")
		arr = append(arr, s)
	}
	return strings.Join(arr, "|")
}
