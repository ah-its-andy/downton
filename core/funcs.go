package core

type Predicate[T any] func(T) bool
type Action[T any] func(T)
type Func[T any, R any] func(T) R
