package enum

type Enum[T comparable] struct {
	members []Member[T]
}

type Member[T comparable] struct {
	value T
}

func M[T comparable](value T) Member[T] {
	return Member[T]{value}
}

func New[T comparable](members ...Member[T]) Enum[T] {
	return Enum[T]{members}
}

func (e Enum[T]) Parse(value T) *Member[T] {
	for _, member := range e.members {
		if member.value == value {
			return &member
		}
	}
	return nil
}

func (e Enum[T]) Value(member Member[T]) T {
	return member.value
}

func (e Enum[T]) Index(member Member[T]) int {
	for i, m := range e.members {
		if m == member {
			return i
		}
	}
	panic("the given Member does not belong to this Enum")
}

func (e Enum[T]) Members() []Member[T] {
	return e.members
}
