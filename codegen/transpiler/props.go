package transpiler

// BoolProps is a map of the boolean properties in a template to their
// vecty names
var BoolProps = map[string]string{
	"autofocus": "Autofocus",
	"checked":   "Checked",
}

// StringProps is a map of the string properties in a template to their
// vecty names
var StringProps = map[string]string{
	"for":         "For",
	"href":        "Href",
	"id":          "ID",
	"placeholder": "Placeholder",
	"src":         "Src",
	"value":       "Value",
}

// TypeProps is a map of the types in a template to their vecty names
var TypeProps = map[string]string{
	"button":         "TypeButton",
	"checkbox":       "TypeCheckbox",
	"color":          "TypeColor",
	"date":           "TypeDate",
	"datetime":       "TypeDatetime",
	"datetime-local": "TypeDatetimeLocal",
	"email":          "TypeEmail",
	"file":           "TypeFile",
	"hidden":         "TypeHidden",
	"image":          "TypeImage",
	"month":          "TypeMonth",
	"number":         "TypeNumber",
	"password":       "TypePassword",
	"radio":          "TypeRadio",
	"range":          "TypeRange",
	"min":            "TypeMin",
	"max":            "TypeMax",
	"value":          "TypeValue",
	"step":           "TypeStep",
	"reset":          "TypeReset",
	"search":         "TypeSearch",
	"submit":         "TypeSubmit",
	"tel":            "TypeTel",
	"text":           "TypeText",
	"time":           "TypeTime",
	"url":            "TypeUrl",
	"week":           "TypeWeek",
}
