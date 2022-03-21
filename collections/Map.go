package collections

type Map[K comparable, V interface{}] interface {
	Iterator[KvPair[K, V]]
	Put(K, V)
	Get(K) V
	Remove(K)
	Contains(K) bool
	Size() int
	Clear()
}

type KvPair[K interface{}, V interface{}] struct {
	Key   K
	Value V
}
