package route

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/factorapp/factor/component"
	"github.com/factorapp/factor/files"
)

// ProcessAll processes components starting at base
func ProcessAll(base string) error {

	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && files.IsHTML(info) {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			//c, _ := component.Parse(f, componentName(path))

			route := files.RouteName(path)
			gfn := filepath.Join(base, strings.ToLower(route)+".go")
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
			transpiler, err := component.NewTranspiler(f, makeStruct, route, "routes")
			if err != nil {
				log.Println("ERROR", err)
				return err
			}

			gofile, err := os.Create(files.GeneratedGoFileName(base, route))
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
