package collections

import (
	"reflect"
	"strings"
)

var StringComparison BinarySearchComparer = func(a, b any) int {
	return strings.Compare(a.(string), b.(string))
}

var IntComparison BinarySearchComparer = func(a, b any) int {
	v := int(a.(int64) - b.(int64))
	if v == 0 {
		return 0
	} else if v < 0 {
		return -1
	} else {
		return 1
	}
}

var FloatComparison BinarySearchComparer = func(a, b any) int {
	v := a.(float64) - b.(float64)
	if v == 0 {
		return 0
	} else if v < 0 {
		return -1
	} else {
		return 1
	}
}

var PtrComparison BinarySearchComparer = func(a, b any) int {
	t := reflect.TypeOf(a)
	if t.Kind() == reflect.Ptr {
		if a == b {
			return 0
		} else {
			return 1
		}
	} else {
		if &a == &b {
			return 0
		} else {
			return 1
		}
	}
}

var BoolComparison BinarySearchComparer = func(a, b any) int {
	if a.(bool) == b.(bool) {
		return 0
	} else {
		return 1
	}
}
