package collections

type List[T any] interface {
	Collection[T]
	Add(T)
	AddRange(...T)
	Remove(T)
	RemoveAt(int)
	Contains(T) bool
	IndexOf(T) int
	Get(int) T
	Set(int, T)
	Size() int
	Clear()
}
