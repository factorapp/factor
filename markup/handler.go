package markup

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/murlokswarm/log"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

// HandleEvent is a helper function to handle events.
// If name designates a component method, the method will be called with argJSON
// unmarshaled into the first arg.
// If name designates a component field, argJSON "Value" field will be directly
// mapped in the component field.
func HandleEvent(nodeID uuid.UUID, name string, argJSON string) {
	if len(name) == 0 {
		log.Panic("no handler")
	}

	n, mounted := nodes[nodeID]
	if !mounted {
		log.Error(errors.Errorf("node with ID = %v does not belong to a mounted component.", nodeID))
		return
	}

	c := n.Mount
	v := reflect.ValueOf(c)

	if m := v.MethodByName(name); m.IsValid() {
		if err := callComponentMethod(m, argJSON); err != nil {
			err = errors.Wrapf(err, "unable to call %v", name)
			log.Error(err)
		}
		return
	}

	pv, err := getPipedValue(v, strings.Split(name, "."))
	if err != nil {
		log.Errorf("unable to map %v: %v", name, err)
		return
	}

	if err = mapPipedValue(pv, argJSON); err != nil {
		log.Errorf("unable to map %v: %v", name, err)
		return
	}
}

func callComponentMethod(m reflect.Value, argJSON string) error {
	t := m.Type()
	numIn := t.NumIn()

	if numIn > 1 {
		return errors.Errorf("func has more than 1 arg: %v", numIn)
	}

	if numIn == 0 {
		m.Call([]reflect.Value{})
		return nil
	}

	argt := t.In(0)
	argv := reflect.New(argt)
	argi := argv.Interface()
	arg := argv.Elem()

	if err := json.Unmarshal([]byte(argJSON), argi); err != nil {
		return errors.Wrapf(err, "unmarshal %v failed", argJSON)
	}

	m.Call([]reflect.Value{arg})
	return nil
}

func getPipedValue(v reflect.Value, pipeline []string) (rv reflect.Value, err error) {
	if len(pipeline) == 0 {
		rv = v
		return
	}
	if len(pipeline[0]) == 0 {
		err = errors.New("pipeline element can't be empty")
		return
	}

	switch k := v.Kind(); k {
	case reflect.Ptr:
		return getPipedPtrValue(v, pipeline)

	case reflect.Struct:
		return getPipedStructFieldValue(v, pipeline)

	case reflect.Map:
		return getPipedMapValue(v, pipeline)

	default:
		err = errors.Errorf("%v is not a valid pipeline source", k)
	}
	return
}

func getPipedPtrValue(v reflect.Value, pipeline []string) (rv reflect.Value, err error) {
	if v.IsNil() {
		t := v.Type()
		nv := reflect.New(t.Elem())
		v.Set(nv)
	}
	return getPipedValue(v.Elem(), pipeline)
}

func getPipedStructFieldValue(v reflect.Value, pipeline []string) (rv reflect.Value, err error) {
	fv := v.FieldByName(pipeline[0])
	if !fv.IsValid() {
		err = errors.Errorf("no field named %v", pipeline[0])
		return
	}
	return getPipedValue(fv, pipeline[1:])
}

func getPipedMapValue(v reflect.Value, pipeline []string) (rv reflect.Value, err error) {
	t := v.Type()
	kt := t.Key()
	if k := kt.Kind(); k != reflect.String {
		err = errors.Errorf("%v key type must be a %v: %v",
			pipeline[0],
			reflect.String,
			k)
		return
	}

	if v.IsNil() {
		nv := reflect.MakeMap(t)
		v.Set(nv)
	}

	kv := reflect.ValueOf(pipeline[0])
	vv := reflect.Zero(t.Elem())
	v.SetMapIndex(kv, vv)
	return getPipedValue(vv, pipeline[1:])
}

func mapPipedValue(v reflect.Value, argJSON string) error {
	var arg struct {
		Value string
	}
	if err := json.Unmarshal([]byte(argJSON), &arg); err != nil {
		return errors.Wrapf(err, "unmarshal %v failed", argJSON)
	}

	if v.Kind() == reflect.String {
		v.SetString(arg.Value)
		return nil
	}

	pv := reflect.New(v.Type())
	pi := pv.Interface()
	if err := json.Unmarshal([]byte(arg.Value), pi); err != nil {
		return errors.Wrapf(err, "unmarshal %v failed", arg.Value)
	}
	v.Set(pv.Elem())
	return nil
}
