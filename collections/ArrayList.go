package collections

var _ List[string] = (*arrayList[string])(nil)

type arrayList[T comparable] struct {
	capacity int
	size     int
	data     []T
}

func NewArrayList[T comparable]() List[T] {
	return &arrayList[T]{
		capacity: 4,
		size:     0,
		data:     make([]T, 4),
	}
}

func NewArrayListWithCapacity[T comparable](capacity int) List[T] {
	return &arrayList[T]{
		capacity: capacity,
		size:     0,
		data:     make([]T, capacity),
	}
}

func NewArrayListWithData[T comparable](data []T) List[T] {
	return &arrayList[T]{
		capacity: len(data),
		size:     len(data),
		data:     data,
	}
}

func (this *arrayList[T]) Add[T](e T) {
	if this.size == this.capacity {
		this.resize(this.capacity * 2)
	}
	this.data[this.size] = e
	this.size++
}

func (this *arrayList[T]) AddAll[T](c Collection[T]) {
	for _, e := range c.ToArray() {
		this.Add(e)
	}
}

func (this *arrayList[T]) Clear[T]() {
	this.size = 0
}

func (this *arrayList[T]) Contains[T](e T) bool {
	for _, o := range this.data {
		if o == e {
			return true
		}
	}
	return false
}

func (this *arrayList[T]) Get[T](index int) T {
	if index < 0 || index >= this.size {
		panic("Index out of bounds")
	}
	return this.data[index]
}

func (this *arrayList[T]) IndexOf[T](e T) int {
	for i, o := range this.data {
		if o == e {
			return i
		}
	}
	return -1
}

func (this *arrayList[T]) IsEmpty[T]() bool {
	return this.size == 0
}

func (this *arrayList[T]) LastIndexOf[T](e T) int {
	for i := this.size - 1;

func (e *arrayList[T]) Add(item T) {
	if e.size == e.capacity {
		e.capacity *= 2
		e.data = append(e.data, make([]T, e.capacity)...)
	}
	e.data[e.size] = item
	e.size++
}

func (e *arrayList[T]) Remove(item T) {
	for i := 0; i < e.size; i++ {
		if e.data[i] == item {
			e.data[i] = e.data[e.size-1]
			e.size--
			return
		}
	}
}

func (e *arrayList[T]) Contains(item T) bool {
	for i := 0; i < e.size; i++ {
		if e.data[i] == item {
			return true
		}
	}
	return false
}

func (e *arrayList[T]) Size() int {
	return e.size
}

func (e *arrayList[T]) Clear() {
	e.size = 0
}

func (e *arrayList[T]) Get(index int) T {
	return e.data[index]
}

func (e *arrayList[T]) Set(index int, item T) {
	e.data[index] = item
}

func (e *arrayList[T]) ToArray() []T {
	ret := make([]T, e.size)
	copy(ret, e.data[:e.size])
	return ret
}
