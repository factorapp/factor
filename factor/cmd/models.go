package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bketelsen/factor/factor/model"
)

func processModels(base string) error {
	return filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !isModel(info) {
			return nil
		}
		model := model.New(path, info)
		toCreate := filepath.Join(base, model.GeneratedFilename())
		log.Printf("generating model %s", toCreate)
		fd, err := os.Create(toCreate)
		if err != nil {
			return err
		}
		if err := model.Write(fd); err != nil {
			return err
		}
		return nil
	})
}

func isModel(info os.FileInfo) bool {
	// ignore generated files
	if strings.HasSuffix(info.Name(), "_generated.go") {
		return false
	}
	// assume all other go files are models
	return filepath.Ext(info.Name()) == ".go"
}
