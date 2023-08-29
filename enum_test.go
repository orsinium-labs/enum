package enum_test

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/enum"
)

type Color = enum.Member[string]

var (
	Red    = Color{"red"}
	Green  = Color{"green"}
	Blue   = Color{"blue"}
	Colors = enum.New[string](Red, Green, Blue)
)

func TestInterfaces(t *testing.T) {
	var _ fmt.Stringer = enum.Member[any]{}
	var _ fmt.GoStringer = enum.Member[any]{}
}

func TestMember_Value(t *testing.T) {
	is := is.New(t)
	is.Equal(Red.Value, "red")
	is.Equal(Green.Value, "green")
	is.Equal(Blue.Value, "blue")
	is.Equal(enum.Member[string]{"blue"}.Value, "blue")
	is.Equal(enum.Member[int]{14}.Value, 14)
}

func TestMember_String(t *testing.T) {
	is := is.New(t)
	act := fmt.Sprint(Red)
	is.Equal(act, "red")
}

func TestMember_GoString(t *testing.T) {
	is := is.New(t)
	act := fmt.Sprintf("%#v", Red)
	is.Equal(act, `enum.Member[string]{"red"}`)
}

func TestEnum_Parse(t *testing.T) {
	is := is.New(t)
	parsed := Colors.Parse("red")
	is.Equal(parsed, &Red)
	parsed = Colors.Parse("purple")
	is.Equal(parsed, nil)
}

func TestEnum_Empty(t *testing.T) {
	is := is.New(t)
	is.True(!Colors.Empty())
	is.True(enum.New[int, enum.Member[int]]().Empty())
}

func TestEnum_Len(t *testing.T) {
	is := is.New(t)
	is.Equal(Colors.Len(), 3)
	is.Equal(enum.New[int, enum.Member[int]]().Len(), 0)
}

func TestEnum_Contains(t *testing.T) {
	is := is.New(t)
	is.True(Colors.Contains(Red))
	is.True(Colors.Contains(Green))
	is.True(Colors.Contains(Blue))
	blue := Color{"blue"}
	is.True(Colors.Contains(blue))
	purple := Color{"purple"}
	is.True(!Colors.Contains(purple))
}

func TestEnum_Members(t *testing.T) {
	is := is.New(t)
	exp := []Color{Red, Green, Blue}
	is.Equal(Colors.Members(), exp)
}

func TestEnum_Values(t *testing.T) {
	is := is.New(t)
	exp := []string{"red", "green", "blue"}
	is.Equal(Colors.Values(), exp)
}

func TestEnum_Value(t *testing.T) {
	is := is.New(t)
	is.Equal(Colors.Value(Red), "red")
}

func TestEnum_Index(t *testing.T) {
	is := is.New(t)
	is.Equal(Colors.Index(Red), 0)
	is.Equal(Colors.Index(Green), 1)
	is.Equal(Colors.Index(Blue), 2)
}

func TestMixMemberTypes(t *testing.T) {
	type Color = enum.Member[string]
	type Country = enum.Member[string]
	red := Color{"red"}
	netherlands := Country{"Netherlands"}
	_ = enum.New[string](red, netherlands)
}
