package markup

import (
	"bytes"
	"encoding/json"
	"text/template"
	"time"
)

// TemplateFuncMapper is the interface that wraps FuncMaps method.
type TemplateFuncMapper interface {
	// Allows to add custom functions to the template used to render the
	// component.
	// Note that funcs named json and time are already implemented to handle
	// structs as prop and time format. Overloads of these will be ignored.
	// See https://golang.org/pkg/text/template/#Template.Funcs for more details.
	FuncMaps() template.FuncMap
}

func render(c Componer) (rendered string, err error) {
	var b bytes.Buffer

	fnmap := template.FuncMap{}
	if t, ok := c.(TemplateFuncMapper); ok {
		extrafuncs := t.FuncMaps()
		for k, v := range extrafuncs {
			fnmap[k] = v
		}
	}
	fnmap["json"] = convertToJSON
	fnmap["time"] = formatTime

	tmpl := template.Must(template.New("Render").Funcs(fnmap).Parse(c.Render()))
	if err = tmpl.Execute(&b, c); err != nil {
		return
	}

	rendered = b.String()
	return
}

func convertToJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return template.HTMLEscapeString(string(b))
}

func formatTime(t time.Time, layout string) string {
	return t.Format(layout)
}
