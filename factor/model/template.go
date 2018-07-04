package model

import (
	"text/template"
)

var modelTypesTpl = template.Must(template.New("modelTypes").Parse(`package models

import (
	"context"
	
	"github.com/satori/go.uuid"
)

type Create{{.UpperName}}Req struct {
	Ctx context.Context
	Data {{.UpperName}}
}

type Create{{.UpperName}}Res struct {
	Err error
}

type Get{{.UpperName}}Req struct {
	ID uuid.UUID
}

type Get{{.UpperName}}Res struct {
	Data {{.UpperName}}
}

type Update{{.UpperName}}Req struct {
	ID uuid.UUID
	New {{.UpperName}}
}

type Update{{.UpperName}}Res struct {
	Err error
}

type Delete{{.UpperName}}Res struct {
	ID uuid.UUID
}

type Delete{{.UpperName}}Res {
	Err error
}`))

var modelServerTpl = template.Must(template.New("model").Parse(`package models

import (
	"net/rpc"
)

{{define "server"}}{{.UpperName}}Server
var {{.LowerName}}Server = New{{.UpperName}}Server()

type {{.UpperName}}Server struct{}

func New{{.UpperName}}Server() {{.UpperName}}Server {
	return &{{.UpperName}}Server{}
}

func (srv *{{.UpperName}}Server) {{.UpperName}}Server) Create(req Create{{.UpperName}}Req, res *Create{{.UpperName}}Res) error {
	// TODO
	return nil
}

func (srv *{{.UpperName}}Server) Get(req Get{{.UpperName}}Req, res *Get{{.UpperName}}Res) error {
	// TODO
	return nil
}

func (srv *{{.UpperName}}Server) Update(req Update{{.UpperName}}Req, res *Update{{.UpperName}}Res) error {
	// TODO
	return nil
}

func (srv *{{.UpperName}}Server) Delete(req Delete{{.UpperName}}Req, res *Delete{{.UpperName}}Res) error {
	// TODO
	return nil
}

func init() {
	rpc.RegisterName("{{.UpperName}}", {{.LowerName}}Server)
}
`))

var modelClientTpl = template.Must(template.New("model").Parse(`package models

import (
	"context"
	"net/rpc"
)

type {{.UpperName}}Client struct{
	RPC *rpc.Client
}

func (cl *{{.UpperName}}Client) Get(id uuid.UUID) (*{{.UpperName}}, error) {
	req := Get{{.UpperName}}Req{ID: id}
	var res Get{{.UpperName}}Res
	if err := cl.RPC.Call("{{.UpperName}}.Get", req, &res); err != nil {
		return nil, err
	}
	return res.Data, nil
}

// TODO: more
`))
