package codegen

import (
	"bytes"
	"text/template"
)

// ServerGoMain returns the main.go for the server
func ServerGoMain(appPath string) (string, error) {
	b := new(bytes.Buffer)
	data := map[string]string{
		"AppPath": appPath,
	}
	if err := serverGoTemplate.Execute(b, data); err != nil {
		return "", err
	}
	return string(b.Bytes()), nil

}

var serverGoTemplate = template.Must(template.New("servergo").Parse(serverGoTemplateStr))

var serverGoTemplateStr = `// build !js,wasm
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"

	{{.AppPath}}/models"
)

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "./app/example.wasm")
}
func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	tds := new(models.TodoServer)
	s.RegisterService(tds, "TodoServer")
	http.HandleFunc("/app/example.wasm", wasmHandler)
	http.Handle("/rpc", s)
	/*	cwd, err := os.Getcwd()
		if err != nil {
			panic(err)
		}
		app := filepath.Join(cwd, "app")
	*/

	http.HandleFunc("/wasm_exec.js", jsHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func jsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./app/wasm_exec.js")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./app/index.html")
}
`
