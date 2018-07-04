package model

import (
	"io"
)

func WriteClientFile(wr io.Writer, names []string) error {
	return clientsTpl.Execute(wr, map[string]interface{}{
		"Clients": names,
	})
}
