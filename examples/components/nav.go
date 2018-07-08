package components

import (
	"github.com/gowasm/vecty"
)

type Nav struct {
	vecty.Core
	MyProp      string `vecty:"Prop"`
	CurrentPath string
}
