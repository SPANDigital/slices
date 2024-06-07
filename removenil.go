package slices

import "reflect"

func isNil[T any](t T) bool {
	v := reflect.ValueOf(t)
	kind := v.Kind()
	// Must be one of these types to be nillable
	return (kind == reflect.Ptr ||
		kind == reflect.Interface ||
		kind == reflect.Slice ||
		kind == reflect.Map ||
		kind == reflect.Chan ||
		kind == reflect.Func) &&
		v.IsNil()
}

func RemoveNil[T any](in []T) (out []T) {
	for _, v := range in {
		if !isNil(v) {
			out = append(out, v)
		}
	}
	return
}
