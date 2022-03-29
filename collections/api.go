package collections

import "sort"

func ForEach[T any, C Iteratable[T]](src C, f func(T)) {
	it := src.GetIterator()
	for it.MoveNext() {
		f(it.Current())
	}
}

func Range[T any](v T, count int) List[T] {
	dest := NewArrayList[T](count)
	for i := 0; i < count; i++ {
		dest.Add(v)
	}
	return dest
}

func Count[T any, C Iteratable[T]](src C, predicate func(T) bool) int {
	var count int
	it := src.GetIterator()
	for it.MoveNext() {
		if predicate(it.Current()) {
			count++
		}
	}
	return count
}

func Distinct[T any, C Collection[T]](src C, comparer BinarySearchComparer) Collection[T] {
	innerComparer := comparer
	if innerComparer == nil {
		innerComparer = func(a1, a2 any) int {
			if a1 == a2 {
				return 0
			} else {
				return -1
			}
		}
	}
	dest := NewArrayList[T](0)
	it := src.GetIterator()
	prev := it.Current()
	for it.MoveNext() {
		if innerComparer(prev, it.Current()) != 0 {
			dest.Add(prev)
		}
		prev = it.Current()
	}
	return dest
}

func OrderBy[T any](src Collection[T], comparer BinarySearchComparer) Collection[T] {
	items := src.ToArray()
	sort.Slice(items, func(i, j int) bool {
		return comparer(items[i], items[j]) < 0
	})
	ret := NewArrayList[T](0)
	ret.AddRange(items...)
	return ret
}
