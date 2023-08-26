package enum

type Enum[M, V comparable] struct {
	members []M
}

func New[V, M comparable](members ...M) Enum[M, V] {
	return Enum[M, V]{members}
}

func (e Enum[M, V]) Parse(value V) (M, bool) {
	for _, m := range e.members {
		if V(m) == value {
			return m, true
		}
	}
	return *new(M), false
}
