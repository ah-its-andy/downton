package collections

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
	dest := NewArrayList[T](0)
	it := src.GetIterator()
	prev := it.Current()
	for it.MoveNext() {
		if comparer(prev, it.Current()) != 0 {
			dest.Add(prev)
		}
		prev = it.Current()
	}
	return dest
}
