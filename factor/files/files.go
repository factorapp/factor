package files

import (
	"path/filepath"
	"regexp"
	"strings"
)

var fileNameRegexp = regexp.MustCompile("[^a-zA-Z0-9]+")

func SanitizedName(path string) string {
	base := filepath.Base(path)
	base = strings.Replace(base, filepath.Ext(path), "", -1)
	return strings.Title(fileNameRegexp.ReplaceAllString(base, ""))
}

func GeneratedName(path string) string {
	n := SanitizedName(path)
	return n + "_generated.go"
}
