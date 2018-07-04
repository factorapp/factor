package model

import (
	"io"
)

func WriteServerFile(wr io.Writer) error {
	return serversTpl.Execute(wr, nil)
}
