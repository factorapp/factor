package components

import (
	"fmt"

	"github.com/bketelsen/factor/golden/models"
	"github.com/bketelsen/factor/markup"
	"github.com/satori/go.uuid"
)

var AppTemplate = `<main>
    <Nav />
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
	tdc := new(models.TodoClient)
	uid := uuid.Must(uuid.NewV4())
	todo, err := tdc.Get(uid)
	t.Todos = []*models.Todo{todo}

	fmt.Println(todo, err)
	return AppTemplate
}

func init() {
	markup.Register(&App{})
}
