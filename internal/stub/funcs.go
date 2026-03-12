package stub

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/Rom4eg/gostub/internal/stub/funcs"
)

func FuncMap() template.FuncMap {
	m := sprig.FuncMap()
	m["sleep"] = funcs.Sleep
	return m
}
