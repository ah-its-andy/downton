package collections

type Map[K comparable, V any] interface {
	Iterator[KvPair[K, V]]
	Put(K, V)
	Get(K) V
	Remove(K)
	Contains(K) bool
	Size() int
	Clear()
}

type KvPair[K any, V any] struct {
	Key   K
	Value V
}
