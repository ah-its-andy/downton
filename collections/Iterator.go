package collections

type Iterator[T any] interface {
	ToArray() []T
}
