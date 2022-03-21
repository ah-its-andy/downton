package collections

import "testing"

func Test_ArrayList_BinarySearch_Validations(t *testing.T) {
	count := 4
	list := NewArrayList[string](count).(*ArrayList[string])
	element := ""
	_, err := list.BinarySearch(0, -1, element, nil)
	if err != ErrBinarySearchLengthNeedNonNegNum {
		t.Error("Expected ErrBinarySearchLengthNeedNonNegNum, but got", err)
	}
	_, err = list.BinarySearch(0, count+1, element, nil)
	if err != ErrBinarySearchInvalidOffLen {
		t.Error("Expected ErrBinarySearchInvalidOffLen, but got", err)
	}
}

func Test_ArrayList_BinarySearch_ForEveryItemWithDuplicates {
	count := 4
	
}