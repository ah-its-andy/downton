package collections

type Set[T comparable] interface {
	Iterator[T]

	Add(T)
	Remove(T)
	Contains(T) bool
	Size() int
	Clear()
	Intersect(Set[T])
	Union(Set[T])
	Difference(Set[T])
	IsSubset(Set[T]) bool
	IsSuperset(Set[T]) bool
	IsDisjoint(Set[T]) bool
}
