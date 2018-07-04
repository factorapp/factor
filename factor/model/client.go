package model

import (
	"io"
)

func WriteClientFile(wr io.Writer) error {
	return clientsTpl.Execute(wr, nil)
}
