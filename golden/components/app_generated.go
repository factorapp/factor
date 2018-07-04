package components

import (
	"fmt"
	"net/url"
	"strings"
	"syscall/js"

	// need the components registered
	"github.com/bketelsen/factor/golden/models"
	_ "github.com/bketelsen/factor/golden/routes"
	"github.com/bketelsen/factor/markup"
	"github.com/satori/go.uuid"
)

var AppTemplate = `<main>
	<Nav />
	
	{{ .Page }}
    {{ range .Todos }}
    <Todo Name="{{.Name}}" Description="{{.Description}}" Permalink="{{.Permalink}}" />
    {{ end }}
    <div>Brian Was Here</div>
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
	fmt.Println(loc)
	u, err := parse(loc.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Path)
	path := cleanPath(u.Path)
	fmt.Println("Path:", path)
	switch cleanPath(path) {
	case "blue":
		t.Page = "<Index />"
		return AppTemplate
	default:

		t.Page = ""
		tdc := new(models.TodoClient)
		uid := uuid.Must(uuid.NewV4())
		todo, err := tdc.Get(uid)
		t.Todos = []*models.Todo{todo}

		fmt.Println(todo, err)
		return AppTemplate

	}
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
