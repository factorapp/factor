package model

import (
	"io"
	"os"
	"strings"

	"github.com/factorapp/factor/files"
)

// Model represents a model
type Model struct {
	fullPath  string
	lowerName string
}

func (m Model) ServerName() string {
	return strings.Title(m.lowerName) + "Server"
}

func (m Model) ClientName() string {
	return strings.Title(m.lowerName) + "Client"
}

func (m Model) TypesFilename() string {
	return m.lowerName + "_types.go"
}

func (m Model) ServerFilename() string {
	return m.lowerName + "_server.go"
}

func (m Model) ClientFilename() string {
	return m.lowerName + "_client.go"
}

func (m Model) Write(typesW io.Writer, clientW io.Writer, serverW io.Writer) error {
	data := map[string]interface{}{
		"UpperName": strings.Title(m.lowerName),
		"LowerName": m.lowerName,
	}
	if err := modelTypesTpl.Execute(typesW, data); err != nil {
		return err
	}
	if err := modelServerTpl.Execute(serverW, data); err != nil {
		return err
	}
	if err := modelClientTpl.Execute(clientW, data); err != nil {
		return err
	}
	return nil
}

func New(fullPath string, info os.FileInfo) Model {
	return Model{
		fullPath:  fullPath,
		lowerName: strings.ToLower(files.SanitizedName(info.Name())),
	}
}
