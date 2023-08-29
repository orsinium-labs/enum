package enum

import (
	"fmt"
)

// Member is an enum member, a specific value bound to a variable.
type Member[T comparable] struct {
	// Value is the underlying of the enum member.
	//
	// You must never change it. It is public only so that
	// it's easier for you to construct new members
	// from the given type alias.
	Value T
}

type iMember[T comparable] interface {
	GetValue() T
}

// GetValue is for internal use only.
//
// If you need to access an enum member value, use Member.Value
// or [Enum.Value] instead.
func (m Member[T]) GetValue() T {
	return m.Value
}

// TypeName is a string representation of the wrapped type.
func (m Member[T]) TypeName() string {
	return fmt.Sprintf("%T", *new(T))
}

// String implements [fmt.Stringer] interface.
//
// It returns the string representation of the wrapped value.
// So, these two lines of code are equivalent:
//
//	fmt.Println(Red)
//	fmt.Println(Red.Value)
func (m Member[T]) String() string {
	return fmt.Sprintf("%v", m.Value)
}

// GoString implements [fmt.GoStringer] interface.
//
// When you print a member using "%#v" format, i
// it will show the member representation as a valid Go syntax.
func (m Member[T]) GoString() string {
	return fmt.Sprintf("%T{%#v}", m, m.Value)
}

// Enum is a collection of enum members.
//
// Use [New] to construct a new Enum from a list of members.
type Enum[M iMember[T], T comparable] struct {
	members []M
}

// New constructs a new [Enum] wrapping the given enum members.
func New[T comparable, M iMember[T]](members ...M) Enum[M, T] {
	return Enum[M, T]{members}
}

// TypeName is a string representation of the wrapped type.
func (m Enum[M, T]) TypeName() string {
	return fmt.Sprintf("%T", *new(T))
}

// Empty returns true if the enum doesn't have any members.
func (e Enum[M, T]) Empty() bool {
	return len(e.members) == 0
}

// Len returns how many mambers the enum has.
func (e Enum[M, T]) Len() int {
	return len(e.members)
}

// Contains returns true if the enum has the given member.
func (e Enum[M, T]) Contains(member M) bool {
	for _, m := range e.members {
		if m.GetValue() == member.GetValue() {
			return true
		}
	}
	return false
}

// Parse converts a raw value into a member of the enum.
//
// If none of the enum members has the given value, nil is returned.
func (e Enum[M, T]) Parse(value T) *M {
	for _, member := range e.members {
		if member.GetValue() == value {
			return &member
		}
	}
	return nil
}

// Value returns Member.Value of the given enum member.
func (e Enum[M, T]) Value(member M) T {
	return member.GetValue()
}

// Index returns the index of the given member in the enum.
//
// If the given memeber is not in the enum, it panics.
func (e Enum[M, T]) Index(member M) int {
	for i, m := range e.members {
		if m.GetValue() == member.GetValue() {
			return i
		}
	}
	panic("the given Member does not belong to this Enum")
}

// Members returns a slice of the members in the enum.
func (e Enum[M, T]) Members() []M {
	return e.members
}

// Values returns a slice of values of all memebers of the enum.
func (e Enum[M, T]) Values() []T {
	res := make([]T, 0, len(e.members))
	for _, m := range e.members {
		res = append(res, m.GetValue())
	}
	return res
}
