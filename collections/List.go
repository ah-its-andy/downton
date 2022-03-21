package collections

type List[T interface{}] interface {
	Iterator[T]
	Add(T)
	Remove(T)
	Contains(T) bool
	Size() int
	Clear()
	GetCapacity() int
}
