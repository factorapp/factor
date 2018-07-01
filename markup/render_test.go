package markup

import (
	"testing"
	"text/template"
	"time"
)

type CompoFuncMapper struct {
	Greet string
}

func (c *CompoFuncMapper) Render() string {
	return `<p>Hello, {{lunny .Greet}}!</p>`
}

func (c *CompoFuncMapper) FuncMaps() template.FuncMap {
	return template.FuncMap{
		"lunny": func(v interface{}) string {
			return "Lunny"
		},
	}
}

func TestConvertToJSON(t *testing.T) {
	t.Log(convertToJSON(42))
}

func TestFormatTime(t *testing.T) {
	t.Log(formatTime(time.Now(), "2006"))
}

func TestTemplateFuncMapper(t *testing.T) {
	c := &CompoFuncMapper{Greet: "Maxence"}
	expected := "<p>Hello, Lunny!</p>"

	r, err := render(c)
	if err != nil {
		t.Fatal(err)
	}
	if r != expected {
		t.Errorf("r should be %v: %v", expected, r)
	}
}
