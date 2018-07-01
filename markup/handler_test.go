package markup

import (
	"reflect"
	"strings"
	"testing"

	"github.com/satori/go.uuid"
)

type FuncArg struct {
	Number int
	String string
}

type HandlerCompo struct {
	String     string
	Struct     FuncArg
	StructPtr  *FuncArg
	Map        map[string]int
	InvalidMap map[int]string

	isCalledWithNoArg     bool
	isCalledWithSingleArg bool
}

func (c *HandlerCompo) HandlerWithoutArg() {
	c.isCalledWithNoArg = true
}

func (c *HandlerCompo) HandlerWitSingleArg(arg FuncArg) {
	c.Struct = arg
	c.isCalledWithSingleArg = true
}

func (c *HandlerCompo) HandlerWitMultipleArg(arg FuncArg, number int) {
}

func (c *HandlerCompo) Render() string {
	return `<h1>Handlers</h1>`
}

func init() {
	Register(&HandlerCompo{})
}

func TestCallComponentMethod(t *testing.T) {
	var err error
	c := &HandlerCompo{}

	// Call without arg.
	cv := reflect.ValueOf(c)
	mv := cv.MethodByName("HandlerWithoutArg")
	if err := callComponentMethod(mv, ""); err != nil {
		t.Fatal(err)
	}
	if !c.isCalledWithNoArg {
		t.Error("HandlerWithoutArg should have been called")
	}

	// Call with single arg.
	arg := `{"Number": 42, "String": "maxoo"}`
	mv = cv.MethodByName("HandlerWitSingleArg")
	if err = callComponentMethod(mv, arg); err != nil {
		t.Fatal(err)
	}
	if !c.isCalledWithSingleArg {
		t.Error("HandlerWitSingleArg should have been called")
	}
	if c.Struct.Number != 42 {
		t.Error("c.Struct.Number should be 42:", c.Struct.Number)
	}
	if c.Struct.String != "maxoo" {
		t.Error("c.Struct.String should be maxoo:", c.Struct.String)
	}

	// Call with bad JSON.
	arg = `{"Number: 42, "String": "maxoo"}`
	if err = callComponentMethod(mv, arg); err == nil {
		t.Error("err should not be nil")
	}
	t.Log(err)

	// Call with multiple args.
	mv = cv.MethodByName("HandlerWitMultipleArg")
	if err = callComponentMethod(mv, arg); err == nil {
		t.Error("err should not be nil")
	}
	t.Log(err)
}

func TestGetPipedValue(t *testing.T) {
	c := &HandlerCompo{
		String: "Hello",
	}
	v := reflect.ValueOf(c).Elem()

	// Direct path.
	fv, err := getPipedValue(v, strings.Split("String", "."))
	if err != nil {
		t.Fatal(err)
	}
	if k := fv.Kind(); k != reflect.String {
		t.Errorf("fv should be a %v: %v", reflect.String, k)
	}

	// Path to struct ptr.
	if fv, err = getPipedValue(v, strings.Split("StructPtr.Number", ".")); err != nil {
		t.Fatal(err)
	}
	if k := fv.Kind(); k != reflect.Int {
		t.Errorf("fv should be a %v: %v", reflect.Int, k)
	}

	// Path to struct.
	if fv, err = getPipedValue(v, strings.Split("Struct.Number", ".")); err != nil {
		t.Fatal(err)
	}
	if k := fv.Kind(); k != reflect.Int {
		t.Errorf("fv should be a %v: %v", reflect.Int, k)
	}

	// Path to map.
	if fv, err = getPipedValue(v, strings.Split("Map.Number", ".")); err != nil {
		t.Fatal(err)
	}
	if k := fv.Kind(); k != reflect.Int {
		t.Errorf("fv should be a %v: %v", reflect.Int, k)
	}
}

func TestGetPipedValueErrors(t *testing.T) {
	c := &HandlerCompo{}
	v := reflect.ValueOf(c).Elem()

	// Invalid path.
	_, err := getPipedValue(v, strings.Split(".Number", "."))
	if err == nil {
		t.Error("err should not be nil")
	}
	t.Log(err)

	// Nonexistent field.
	if _, err = getPipedValue(v, strings.Split("Foo", ".")); err == nil {
		t.Error("err should not be nil")
	}
	t.Log(err)

	// Path to invalid map.
	if _, err = getPipedValue(v, strings.Split("InvalidMap.1", ".")); err == nil {
		t.Error("err should not be nil")
	}
	t.Log(err)

	// Invalid pipeline source.
	if _, err = getPipedValue(v, strings.Split("String.max", ".")); err == nil {
		t.Error("err should not be nil")
	}
	t.Log(err)
}

func TestMapPipedValue(t *testing.T) {
	// String.
	arg := `{"Value": "Martine est sanglante"}`
	v := reflect.New(reflect.TypeOf("sd")).Elem()
	if err := mapPipedValue(v, arg); err != nil {
		t.Fatal(err)
	}
	if str := v.Interface().(string); str != "Martine est sanglante" {
		t.Errorf(`str should be "Martine est sanglante": %v`, str)
	}

	// Number.
	arg = `{"Value": "42"}`
	v = reflect.New(reflect.TypeOf(42)).Elem()
	if err := mapPipedValue(v, arg); err != nil {
		t.Fatal(err)
	}
	if nb := v.Interface().(int); nb != 42 {
		t.Errorf(`nb should be 42: %v`, nb)
	}
}

func TestMapPipedValueErrors(t *testing.T) {
	// String.
	arg := `{Value": "Martine est sanglante"}`
	v := reflect.New(reflect.TypeOf("")).Elem()
	err := mapPipedValue(v, arg)
	if err == nil {
		t.Error("err should not be nil")
	}
	t.Log(err)

	// Number.
	arg = `{"Value": "Martine est sanglante"}`
	v = reflect.New(reflect.TypeOf(42)).Elem()
	if err = mapPipedValue(v, arg); err == nil {
		t.Error("err should not be nil")
	}
	t.Log(err)
}

func TestHandleEvent(t *testing.T) {
	c := &HandlerCompo{}
	ctx := uuid.Must(uuid.NewV1())

	root, err := Mount(c, ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer Dismount(c)

	// Method.
	HandleEvent(root.ID, "HandlerWithoutArg", "")
	if !c.isCalledWithNoArg {
		t.Error("HandlerWithoutArg should have been called")
	}

	// Field.
	HandleEvent(root.ID, "String", `{"Value": "hello"}`)
	if c.String != "hello" {
		t.Error("c.String should be hello:", c.String)
	}

	HandleEvent(root.ID, "Struct.Number", `{"Value": "42"}`)
	if c.Struct.Number != 42 {
		t.Error("c.Struct.Number should be 42:", c.Struct.Number)
	}

	HandleEvent(root.ID, "StructPtr.Number", `{"Value": "42"}`)
	if c.StructPtr.Number != 42 {
		t.Error("c.StructPtr.Number should be 42:", c.StructPtr.Number)
	}

	HandleEvent(root.ID, "Struct", `{"Value": "{\"Number\":21}"}`)
	if c.Struct.Number != 21 {
		t.Error("c.Struct.Number should be 21:", c.Struct.Number)
	}
}

func TestHandleEventPanic(t *testing.T) {
	defer func() { recover() }()

	c := &HandlerCompo{}
	ctx := uuid.Must(uuid.NewV1())

	root, err := Mount(c, ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer Dismount(c)

	// Method.
	HandleEvent(root.ID, "", "")
	t.Error("should have panic")
}

func TestHandleEventError(t *testing.T) {
	c := &HandlerCompo{}
	ctx := uuid.Must(uuid.NewV1())

	// Not mounted.
	HandleEvent(uuid.Must(uuid.NewV1()), "HandlerWitMultipleArg", "")

	root, err := Mount(c, ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer Dismount(c)

	// Method.
	HandleEvent(root.ID, "HandlerWitMultipleArg", "")

	// Field.
	HandleEvent(root.ID, "Hello", `{"Value": "hello"}`)
	HandleEvent(root.ID, "String", `{"Value": hello"}`)
}
