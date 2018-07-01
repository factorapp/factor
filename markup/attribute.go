package markup

import (
	"encoding/json"
	"reflect"
	"strconv"
)

// AttributeMap represents a map of attributes.
type AttributeMap map[string]string

func (m AttributeMap) diff(other AttributeMap) AttributeMap {
	res := AttributeMap{}

	for n, v := range m {
		if len(other) == 0 {
			res[n] = ""
			continue
		}

		nv, ok := other[n]
		if !ok {
			res[n] = ""
			continue
		}

		if v != nv {
			res[n] = nv
			continue
		}
	}

	for n, v := range other {
		if len(m) != 0 {
			ov, ok := m[n]
			if ok && ov == v {
				continue
			}
		}

		res[n] = v
	}
	return res
}

func decodeAttributeMap(attributes AttributeMap, c Componer) {
	v := reflect.ValueOf(c).Elem()

	for name, value := range attributes {
		if f := v.FieldByName(name); f.IsValid() {
			decodeValue(f, value)
		}
	}
}

func decodeValue(v reflect.Value, value string) {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Bool:
		b, _ := strconv.ParseBool(value)
		v.SetBool(b)

	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		n, _ := strconv.ParseInt(value, 0, 64)
		v.SetInt(n)

	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uintptr:
		n, _ := strconv.ParseUint(value, 0, 64)
		v.SetUint(n)

	case reflect.Float64, reflect.Float32:
		n, _ := strconv.ParseFloat(value, 64)
		v.SetFloat(n)

	case reflect.Struct, reflect.Slice, reflect.Map:
		s := reflect.New(v.Type())
		json.Unmarshal([]byte(value), s.Interface())
		v.Set(s.Elem())
	}
}
