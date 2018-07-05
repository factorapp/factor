package routes

import (
	"fmt"

	"github.com/bketelsen/factor/golden/models"
	"github.com/bketelsen/factor/markup"
	"github.com/satori/go.uuid"
)

type Todos struct {
	List []*models.Todo
}

//    <Todo Name="{{.Name}}" Description="{{.Description}}" Permalink="{{.Permalink}}"/>
var TodosTemplate = `<div class="todos">
{{ range .List }}
<Todo Name="{{.Name}}" Description="{{.Description}}" Permalink="{{.Permalink}}"/>
{{ end }}
</div>`
var TodosStyles = ``

func (t *Todos) Render() string {
	tdc := new(models.TodoClient)
	uid := uuid.Must(uuid.NewV4())
	fmt.Println(uid.String())
	todo, err := tdc.Get(uid)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	/*	todo := &models.Todo{
			Name:        "Brian",
			Description: "Des",
			Permalink:   "/1",
		}
	*/
	t.List = []*models.Todo{todo}
	return TodosTemplate
}
func (t *Todos) Style() string {
	return TodosStyles
}
func (t *Todos) OnMount() {

}

func init() {
	markup.Register(&Todos{})
}
