package model

import (
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/bketelsen/factor/factor/files"
)

// Model represents a model
type Model struct {
	fullPath  string
	lowerName string
}

func (m Model) GeneratedFilename() string {
	return m.lowerName + "_generated.go"
}

func (m Model) Write(out io.Writer) error {
	data := map[string]interface{}{
		"UpperName": strings.Title(m.lowerName),
	}
	return modelTpl.Execute(out, data)
}

func New(fullPath string, info os.FileInfo) Model {
	return Model{
		fullPath:  fullPath,
		lowerName: strings.ToLower(files.SanitizedName(info.Name())),
	}
}

var modelTpl = template.Must(template.New("model").Parse(`package models

import (
	"context"
)

type {{.UpperName}}Client struct{}

func New{{.UpperName}}Client() {{.UpperName}}Client {
	return &{{.UpperName}}Client{}
}

type New{{.UpperName}} struct {
	Ctx context.Context
	Data {{.UpperName}}
}

type {{.UpperName}}Resp struct {
	Err error
}`))
