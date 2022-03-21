package collections

var _ List[string] = (*ArrayList[string])(nil)

type ArrayList[T comparable] struct {
	capacity int
	size     int
	data     []T
	Comparer BinarySearchComparer[T]
}

func (c *ArrayList[T]) BinarySearch(startIndex int, offsets int, element T, comparer BinarySearchComparer[T]) int {
	low := 0
	high := c.size - 1
	for low <= high {
		mid := (low + high) / 2
		cur := c.data[mid]
		if !comparer(cur, element) {
			low = mid + 1
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
