package route

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/factorapp/factor/component"
)

// ProcessAll processes components starting at base
func ProcessAll(base string) error {

	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isHTML(info) {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			//c, _ := component.Parse(f, componentName(path))

			comp := componentName(path)
			gfn := filepath.Join(base, strings.ToLower(comp)+".go")
			_, err = os.Stat(gfn)
			var makeStruct bool
			if os.IsNotExist(err) {
				makeStruct = true
			}
			/*gofile, err := os.Create(goFileName(base, componentName(path)))
			if err != nil {
				return err
			}
			defer gofile.Close()

			c.Transform(gofile)
			*/
			transpiler, err := component.NewTranspiler(f, makeStruct, comp, "routes")
			if err != nil {
				log.Println("ERROR", err)
				return err
			}

			gofile, err := os.Create(goFileName(base, comp))
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
			defer gofile.Close()
			_, err = io.WriteString(gofile, transpiler.Code())
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("error walking the path %q: %v\n", base, err)
	}
	return err
}

func isHTML(info os.FileInfo) bool {
	return filepath.Ext(info.Name()) == ".html"
}

func goFileName(base, comp string) string {
	return filepath.Join(base, strings.ToLower(comp)+"_generated.go")
}
func componentName(path string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	base := filepath.Base(path)
	base = strings.Replace(base, filepath.Ext(path), "", -1)
	return strings.Title(reg.ReplaceAllString(base, ""))
}
