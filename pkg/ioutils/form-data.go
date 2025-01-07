package ioutils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func ReadFormToStruct[T any](form *multipart.Form, s *T) error {
	v := reflect.Indirect(reflect.ValueOf(s))
	return readMapToStruct(form.Value, v)
}

func readMapToStruct(m map[string][]string, v reflect.Value) (err error) {
	if v.Kind() != reflect.Struct {
		return errors.New("value is not a struct")
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		if !f.IsExported() {
			continue
		}
		key, ok := f.Tag.Lookup("json")
		if !ok {
			key = f.Name
		}
		if key == "-" {
			continue
		}
		val, ok := m[key]
		if ok && len(val) > 0 {
			if f.Type.Kind() == reflect.Slice {
				v.Field(i).Set(mapSliceToType(f.Type, val))
			} else {
				v.Field(i).Set(convertToType(f.Type, val[0]))
			}
			continue
		}
		if m, ok := collectFormNestedFields(m, key); ok {
			readMapToStruct(m, v.Field(i))
		}
	}

	return nil
}

func convertToType(t reflect.Type, v string) reflect.Value {
	// TODO: here must be many more cases, but I'm not ready to implement them all
	switch t.Kind() {
	case reflect.Int:
		n, _ := strconv.ParseInt(v, 10, 0)
		return reflect.ValueOf(n).Convert(t)
	case reflect.Float64:
		r, _ := strconv.ParseFloat(v, 64)
		return reflect.ValueOf(r).Convert(t)
	case reflect.Bool:
		isTruthy := slices.Contains([]string{"1", "true", "yes", "on"}, strings.ToLower(v))
		return reflect.ValueOf(isTruthy).Convert(t)
	case reflect.Slice:
	default:
		return reflect.ValueOf(v)
	}
	return reflect.ValueOf(v)
}

func mapSliceToType(t reflect.Type, vals []string) reflect.Value {
	values := reflect.MakeSlice(t, 0, len(vals))
	for _, e := range vals {
		values = reflect.Append(values, convertToType(t.Elem(), e))
	}
	return values
}

func collectFormNestedFields(m map[string][]string, path string) (map[string][]string, bool) {
	r := make(map[string][]string)
	for k, v := range m {
		if !strings.HasPrefix(k, path+"[") {
			continue
		}
		parts := strings.SplitN(k, "[", 2)
		if len(parts) < 2 {
			continue
		}
		nestedKey := strings.Replace(parts[1], "]", "", 1)
		r[nestedKey] = v
	}

	return r, len(r) > 0
}
