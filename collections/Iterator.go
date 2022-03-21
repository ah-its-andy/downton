package collections

type Iterator[T interface{}] interface {
	ToArray() []T
}
