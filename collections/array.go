package collections

import (
	"errors"
	"reflect"

	"github.com/ah-its-andy/downton/core"
)

type BinarySearchComparer func(any, any) int

var ErrBinarySearchArrayNil = errors.New("Array is nil")
var ErrBinarySearchLengthNeedNonNegNum = errors.New("Length need non-negative number")
var ErrBinarySearchInvalidOffLen = errors.New("Invalid offset or length")

func BinarySearch(array []any, index int, length int, value any, comparer BinarySearchComparer) (int, error) {
	if array == nil {
		return -1, ErrBinarySearchArrayNil
	}
	if length < 0 {
		return -1, ErrBinarySearchLengthNeedNonNegNum
	}
	if len(array)-index < length {
		return -1, ErrBinarySearchInvalidOffLen
	}

	innerComparer := comparer
	if innerComparer == nil {
		itemType := reflect.TypeOf(value)
		if itemType.Kind() == reflect.Ptr {
			itemType = itemType.Elem()
		}

		if itemType.Kind() == reflect.Float32 || itemType.Kind() == reflect.Float64 {
			//todo: call float64 comparer
		} else if itemType.Kind() == reflect.Int ||
			itemType.Kind() == reflect.Int8 ||
			itemType.Kind() == reflect.Int16 ||
			itemType.Kind() == reflect.Int32 ||
			itemType.Kind() == reflect.Int64 {
			//todo: call int64 comparer
		} else if itemType.Kind() == reflect.String {
			//todo: call string comparer
		} else if itemType.Kind() == reflect.Bool {
			//todo: call bool comparer
		}
	}
	lo := index
	hi := index + length - 1
	for lo <= hi {
		middle := (lo + hi) / 2
		cur := array[middle]
		if innerComparer == nil {
			c := innerComparer(cur, value)
			if c == 0 {
				return middle, nil
			} else if c < 0 {
				lo = lo + 1
			} else {
				hi = hi - 1
			}
		} else {
			if comp, ok := value.(core.Comparable); ok {
				c := comp.CompareTo(cur)
				if c == 0 {
					return middle, nil
				} else {
					lo = lo + 1
					hi = hi - 1
				}
			} else {
				if cur == value {
					return middle, nil
				} else {
					lo = lo + 1
					hi = hi - 1
				}
			}
		}
	}
	return -1, nil
}
