package enum_test

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/enum"
)

type Color string

var (
	Red    = enum.M[Color]("red")
	Green  = enum.M[Color]("green")
	Blue   = enum.M[Color]("blue")
	Colors = enum.New(Red, Green, Blue)
)

func TestInterfaces(t *testing.T) {
	var _ fmt.Stringer = enum.Member[any]{}
	var _ fmt.GoStringer = enum.Member[any]{}
}

func TestMemberValue(t *testing.T) {
	is := is.New(t)
	is.Equal(Red.Value(), Color("red"))
	is.Equal(Green.Value(), Color("green"))
	is.Equal(Blue.Value(), Color("blue"))
	is.Equal(enum.M[string]("blue").Value(), "blue")
	is.Equal(enum.M[int](14).Value(), 14)
}

func TestMemberString(t *testing.T) {
	is := is.New(t)
	act := fmt.Sprint(Red)
	is.Equal(act, "red")
}

func TestMemberGoString(t *testing.T) {
	is := is.New(t)
	act := fmt.Sprintf("%#v", Red)
	is.Equal(act, `enum.M[Color]("red")`)
}

func TestEnumParse(t *testing.T) {
	is := is.New(t)
	parsed := Colors.Parse("red")
	is.Equal(parsed, &Red)
}

func TestEnumEmpty(t *testing.T) {
	is := is.New(t)
	is.True(!Colors.Empty())
	is.True(enum.New[int]().Empty())
}

func TestEnumLen(t *testing.T) {
	is := is.New(t)
	is.Equal(Colors.Len(), 3)
	is.Equal(enum.New[int]().Len(), 0)
}

func TestEnumContains(t *testing.T) {
	is := is.New(t)
	is.True(Colors.Contains(Red))
	is.True(Colors.Contains(Green))
	is.True(Colors.Contains(Blue))
	blue := enum.M[Color]("blue")
	is.True(Colors.Contains(blue))
	purple := enum.M[Color]("purple")
	is.True(!Colors.Contains(purple))
}

func TestEnumMembers(t *testing.T) {
	is := is.New(t)
	exp := []enum.Member[Color]{Red, Green, Blue}
	is.Equal(Colors.Members(), exp)
}

func TestEnumValues(t *testing.T) {
	is := is.New(t)
	exp := []Color{"red", "green", "blue"}
	is.Equal(Colors.Values(), exp)
}
