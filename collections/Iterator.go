package collections

type Iterator[T any] interface {
	MoveNext() bool
	Reset()
	Current() *T
}
