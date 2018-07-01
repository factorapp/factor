package markup

import "testing"

type StructProp struct {
	Value int
}

type PropsTest struct {
	String  string
	Bool    bool
	Int     int
	Uint    uint
	Float   float64
	Struct  StructProp
	Slice   []int
	Map     map[string]string
	Unknown int
}

func (p *PropsTest) Render() string {
	return `
<div>
	<p>String: {{.String}}</p>
	<p>Bool: {{.Bool}}</p>
	<p>Int: {{.Int}}</p>
	<p>Uint: {{.Uint}}</p>
	<p>Float: {{.Float}}</p>
    <p>Struct: {{.Struct}}</p>
    <p>Slice: {{.Slice}}</p>
    <p>Map: {{.Map}}</p>
</div>
	`
}

func TestAttributeMapDiff(t *testing.T) {
	m1 := AttributeMap{
		"delete": "bar",
		"same":   "miam",
		"update": "foo",
	}
	m2 := AttributeMap{
		"update": "boo",
		"same":   "miam",
		"add":    "bar",
	}
	res := m1.diff(m2)

	if len(res["delete"]) != 0 {
		t.Error(`res["delete"] should be empty:`, res["delete"])
	}

	if v, ok := res["same"]; ok {
		t.Error(`res["same"] should not be in res:`, v)
	}

	if res["update"] != "boo" {
		t.Error(`res[upate] should be boo:`, res["update"])
	}

	if res["add"] != "bar" {
		t.Error(`res[add] should be bar:`, res["add"])
	}
}

func TestAttributeMapNoDiff(t *testing.T) {
	m1 := AttributeMap{
		"same":  "miam",
		"same2": "boo",
	}

	m2 := AttributeMap{
		"same":  "miam",
		"same2": "boo",
	}

	if l := len(m1.diff(m2)); l != 0 {
		t.Error("l should be 0:", l)
	}
}

func TestAttributeMapDiffOtherEmpty(t *testing.T) {
	m1 := AttributeMap{
		"delete": "dirty",
	}
	m2 := AttributeMap{}
	res := m1.diff(m2)

	if l := len(m1.diff(m2)); l != 1 {
		t.Error("l should be 1:", l)
	}

	if len(res["delete"]) != 0 {
		t.Error(`res[delete] should be empty:`, res["delete"])
	}
}

func TestDecodeAttributeMap(t *testing.T) {
	s := &PropsTest{}
	attributes := AttributeMap{
		"String": "Hello",
		"Bool":   "true",
		"Int":    "-42",
		"Uint":   "42",
		"Float":  "3.14",
		"Struct": `{"Value": 21}`,
		"Slice":  "[21, 42]",
		"Map":    `{"name": "go", "version":"1.42"}`,
	}
	decodeAttributeMap(attributes, s)

	if s.String != "Hello" {
		t.Error("s.String should be Hello:", s.String)
	}

	if !s.Bool {
		t.Error("s.Bool should be true:", s.Bool)
	}

	if s.Int != -42 {
		t.Error("s.Int should be -42:", s.Int)
	}

	if s.Uint != 42 {
		t.Error("s.Uint should be 42:", s.Uint)
	}

	if s.Float != 3.14 {
		t.Error("s.Float should be 3.14:", s.Float)
	}

	if s.Struct.Value != 21 {
		t.Error("s.Struct.Value should be 21:", s.Struct.Value)
	}

	if l := len(s.Slice); l != 2 {
		t.Error("len of s.Slice should be 2:", l)
	}

	if s.Slice[0] != 21 {
		t.Error("s.Slice[0] should be 21:", s.Slice[0])
	}

	if s.Slice[1] != 42 {
		t.Error("s.Slice[1] should be 42:", s.Slice[1])
	}

	if s.Map["name"] != "go" {
		t.Error(`s.Map["name"] should be go:`, s.Map["name"])
	}

	if s.Map["version"] != "1.42" {
		t.Error(`s.Map["version"] should be 1.42:`, s.Map["version"])
	}

	if s.Unknown != 0 {
		t.Error("s.Unknown should ve 0:", s.Unknown)
	}
}
