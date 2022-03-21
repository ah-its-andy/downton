package collections

type List[T comparable] interface {
	Iterator[T]
	Add(T)
	Remove(T)
	Contains(T) bool
	Size() int
	Clear()
	Get(int) T
	Set(int, T)
}
