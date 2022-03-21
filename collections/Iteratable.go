package collections

type Iteratable[T any] interface {
	GetIterator() Iterator[T]
}
