package enum

import (
	"fmt"
	"strings"
)

type Member[T comparable] struct {
	value T
}

func M[T comparable](value T) Member[T] {
	return Member[T]{value}
}

func (m Member[T]) Value() T {
	return m.value
}

func (m Member[T]) TypeName() string {
	// The %T format returns the type name containing
	// the module name in which it is defined.
	// However, enum members are usually defined
	// right next to the type they wrap,
	// so we can drop the package name.
	name := fmt.Sprintf("%T", *new(T))
	parts := strings.Split(name, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return parts[0]
}

func (m Member[T]) String() string {
	return fmt.Sprintf("%v", m.value)
}

func (m Member[T]) GoString() string {
	return fmt.Sprintf("enum.M[%s](%#v)", m.TypeName(), m.value)
}

type Enum[T comparable] struct {
	members []Member[T]
}

func New[T comparable](members ...Member[T]) Enum[T] {
	return Enum[T]{members}
}

func (m Enum[T]) TypeName() string {
	return Member[T]{}.TypeName()
}

func (e Enum[T]) Empty() bool {
	return len(e.members) == 0
}

func (e Enum[T]) Len() int {
	return len(e.members)
}

func (e Enum[T]) Contains(member Member[T]) bool {
	for _, m := range e.members {
		if m == member {
			return true
		}
	}
	return false
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

func (e Enum[T]) Values() []T {
	res := make([]T, len(e.members))
	for _, m := range e.members {
		res = append(res, m.value)
	}
	return res
}
