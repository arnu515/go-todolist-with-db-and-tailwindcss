package util

import "reflect"

func Default[T any](value T, _default T) T {
	if reflect.ValueOf(value).IsZero() {
		return _default
	}
	return value
}
