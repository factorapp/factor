package components

import (
	"fmt"
	"net/url"
	"strings"
	"syscall/js"

	// need the components registered

	"github.com/bketelsen/factor/markup"
)

var AppTemplate = `<main>
	<Nav />
	
	{{ .Page }}

</main>`
var AppStyles = `
    main {
        position: relative;
        max-width: 56em;
        background-color: white;
        padding: 2em;
        margin: 0 auto;
        box-sizing: border-box;
    }
`

func (t *App) Render() string {
	loc := js.Global().Get("location")
	u, err := parse(loc.String())
	if err != nil {
		panic(err)
	}
	path := cleanPath(u.Path)
	t.Page = getPage(path)
	return AppTemplate

}

// should return a tag
func getPage(path string) string {
	if path == "/" {
		return ""
	}
	if path == "" {
		path = "index"
	}
	return fmt.Sprintf("<%s />", strings.Title(path))
}
func init() {
	markup.Register(&App{})
}

func parse(location string) (*url.URL, error) {
	return url.Parse(location)
}

func cleanPath(path string) string {
	return strings.Replace(path, "/", "", -1)
}
