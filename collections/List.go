package collections

type List[T comparable] interface {
	Iterator[T]
	Add(T)
	Remove(T)
	RemoveAt(int)
	Contains(T) bool
	IndexOf(T) int
	Get(int) T
	Set(int, T)
	Size() int
	Clear()
}
