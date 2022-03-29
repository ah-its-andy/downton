package collections

type Collection[T any] interface {
	Iteratable[T]

	ToList() List[T]
	ToArray() []T
	Values() []T
}
