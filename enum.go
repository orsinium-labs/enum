package enum

import (
	"fmt"
	"strings"
)

type Member[T comparable] struct {
	Value T
}

type iMember[T comparable] interface {
	GetValue() T
}

func (m Member[T]) GetValue() T {
	return m.Value
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
	return fmt.Sprintf("%v", m.Value)
}

func (m Member[T]) GoString() string {
	return fmt.Sprintf("enum.M[%s](%#v)", m.TypeName(), m.Value)
}

type Enum[M iMember[T], T comparable] struct {
	members []M
}

func New[T comparable, M iMember[T]](members ...M) Enum[M, T] {
	return Enum[M, T]{members}
}

func (m Enum[M, T]) TypeName() string {
	return Member[T]{}.TypeName()
}

func (e Enum[M, T]) Empty() bool {
	return len(e.members) == 0
}

func (e Enum[M, T]) Len() int {
	return len(e.members)
}

func (e Enum[M, T]) Contains(member M) bool {
	for _, m := range e.members {
		if m.GetValue() == member.GetValue() {
			return true
		}
	}
	return false
}

func (e Enum[M, T]) Parse(value T) *M {
	for _, member := range e.members {
		if member.GetValue() == value {
			return &member
		}
	}
	return nil
}

func (e Enum[M, T]) Value(member M) T {
	return member.GetValue()
}

func (e Enum[M, T]) Index(member M) int {
	for i, m := range e.members {
		if m.GetValue() == member.GetValue() {
			return i
		}
	}
	panic("the given Member does not belong to this Enum")
}

func (e Enum[M, T]) Members() []M {
	return e.members
}

func (e Enum[M, T]) Values() []T {
	res := make([]T, 0, len(e.members))
	for _, m := range e.members {
		res = append(res, m.GetValue())
	}
	return res
}
