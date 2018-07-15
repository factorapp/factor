package files

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var reg = regexp.MustCompile("[^a-zA-Z0-9]+")

func IsHTML(info os.FileInfo) bool {
	return filepath.Ext(info.Name()) == ".html" || filepath.Ext(info.Name()) == ".ghtml"
}

func GeneratedGoFileName(base, name string) string {
	return filepath.Join(base, strings.ToLower(name)+"_generated.go")
}
func ComponentName(path string) string {
	base := filepath.Base(path)
	base = strings.Replace(base, filepath.Ext(path), "", -1)
	return strings.Title(reg.ReplaceAllString(base, ""))
}

func RouteName(path string) string {
	return ComponentName(path)
}
