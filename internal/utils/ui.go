package utils

import "github.com/jedib0t/go-pretty/v6/text"

func ColorString(str string) string {
	return text.FgCyan.Sprint(str)
}
