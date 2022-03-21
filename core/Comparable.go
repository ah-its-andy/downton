package core

type Comparable[T any] interface {
	CompareTo(T) int
}
