package util

import "strings"

func Capitalize(s string) string {
	var ns string

	ws := strings.Split(strings.TrimSpace(s), " ")

	for i := range len(ws) {
		ws[i] = strings.ToUpper(string(ws[i][0])) + ws[i][1:]
	}

	ns = strings.Join(ws, " ")

	return ns
}
