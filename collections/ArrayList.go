package collections

var _ List[string] = (*ArrayList[string])(nil)

func NewArrayList[T any](capacity int) List[T] {
	return &ArrayList[T]{
		capacity: capacity,
		size:     0,
		data:     make([]interface{}, capacity),
	}
}

type ArrayList[T any] struct {
	capacity int
	size     int
	data     []interface{}
	Comparer BinarySearchComparer
}

type arrayListIterator[T any] struct {
	list   *ArrayList[T]
	curPos int
}

func (iterator *arrayListIterator[T]) MoveNext() bool {
	if iterator.curPos < iterator.list.size {
		iterator.curPos++
		return true
	}
	return false
}

func (iterator *arrayListIterator[T]) Current() *T {
	return iterator.list.data[iterator.curPos].(*T)
}

func (iterator *arrayListIterator[T]) Reset() {
	iterator.curPos = 0
}

func (c *ArrayList[T]) BinarySearch(index int, length int, value interface{}, comparer BinarySearchComparer) (int, error) {
	return BinarySearch(c.data, index, length, value, comparer)
}

func (c *ArrayList[T]) increaseCapacity() {
	if c.capacity == 0 {
		c.capacity = 4
	} else {
		c.capacity = c.capacity * 2
	}
	c.data = append(c.data, make([]interface{}, c.capacity)...)
}

func (c *ArrayList[T]) Add(item *T) {
	if c.size == c.capacity {
		c.increaseCapacity()
	}
	c.data[c.size] = item
	c.size++
}

func (c *ArrayList[T]) RemoveAt(itemIndex int) {
	if itemIndex == -1 {
		return
	}
	for i := itemIndex; i < c.size-1; i++ {
		if i == c.size-1 {
			c.data[i] = nil
		} else {
			c.data[i] = c.data[i+1]
		}
	}
	c.size--
}

func (c *ArrayList[T]) Remove(item *T) {
	itemIndex, err := c.BinarySearch(0, c.size, item, c.Comparer)
	if err != nil {
		panic(err)
	}
	c.RemoveAt(itemIndex)
}

func (c *ArrayList[T]) Contains(item *T) bool {
	itemIndex, err := c.BinarySearch(0, c.size, item, c.Comparer)
	if err != nil {
		panic(err)
	}
	return itemIndex != -1
}

func (c *ArrayList[T]) IndexOf(item *T) int {
	itemIndex, err := c.BinarySearch(0, c.size, item, c.Comparer)
	if err != nil {
		panic(err)
	}
	return itemIndex
}

func (c *ArrayList[T]) Get(i int) *T {
	return c.data[i].(*T)
}

func (c *ArrayList[T]) Set(i int, item *T) {
	c.data[i] = item
}

func (c *ArrayList[T]) Size() int {
	return c.size
}

func (c *ArrayList[T]) Clear() {
	c.capacity = 4
	c.data = make([]interface{}, c.capacity)
	c.size = 0
}

func (c *ArrayList[T]) GetIterator() Iterator[T] {
	return &arrayListIterator[T]{
		list:   c,
		curPos: 0,
	}
}

func (c *ArrayList[T]) ToArray() []*T {
	if c.size == 0 {
		return []*T{}
	}
	dest := make([]*T, c.size)

	for i := 0; i < c.size; i++ {
		dest[i] = c.data[i].(*T)
	}
	return dest
}

func (c *ArrayList[T]) ToList() List[T] {
	dest := &ArrayList[T]{
		size:     c.size,
		capacity: c.capacity,
		data:     make([]interface{}, c.capacity),
	}
	copy(dest.data, c.data)
	return dest
}
