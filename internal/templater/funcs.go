package templater

import (
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
}
