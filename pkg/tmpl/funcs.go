package tmpl

import (
	"strings"
	"text/template"
)

var FuncMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
}
