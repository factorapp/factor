package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/factorapp/factor/factor/model"
)

func processModels(base string) error {
	servers := []string{}
	clients := []string{}
	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
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

		typesFilename := filepath.Join(base, model.TypesFilename())
		log.Printf("generating %s", typesFilename)
		typesFd, err := os.Create(typesFilename)
		if err != nil {
			return err
		}

		serverFilename := filepath.Join(base, model.ServerFilename())
		log.Printf("generating %s", serverFilename)
		serverFd, err := os.Create(serverFilename)
		if err != nil {
			return err
		}

		clientFilename := filepath.Join(base, model.ClientFilename())
		log.Printf("generating %s", clientFilename)
		clientFd, err := os.Create(clientFilename)
		if err != nil {
			return err
		}

		if err := model.Write(typesFd, clientFd, serverFd); err != nil {
			return err
		}
		servers = append(servers, model.ServerName())
		clients = append(clients, model.ClientName())
		return nil
	})
	if err != nil {
		return err
	}

	log.Printf("writing client file")
	clientFd, err := os.Create(filepath.Join(base, "client_generated.go"))
	if err != nil {
		return err
	}
	if err := model.WriteClientFile(clientFd, clients); err != nil {
		return err
	}

	log.Printf("writing server file")
	serverFd, err := os.Create(filepath.Join(base, "server_generated.go"))
	if err != nil {
		return err
	}
	if err := model.WriteServerFile(serverFd); err != nil {
		return err
	}
	log.Printf("generated servers and clients: %s", servers)
	return nil
}

func isModel(info os.FileInfo) bool {
	// ignore generated files
	if strings.HasSuffix(info.Name(), "_generated.go") ||
		strings.HasSuffix(info.Name(), "_server.go") ||
		strings.HasSuffix(info.Name(), "_types.go") ||
		strings.HasSuffix(info.Name(), "_client.go") {
		return false
	}
	// assume all other go files are models
	return filepath.Ext(info.Name()) == ".go"
}
