package enum_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/enum"
)

func TestEnumSmoke(t *testing.T) {
	is := is.New(t)
	type Color string
	var (
		Red    = enum.M[Color]("red")
		Green  = enum.M[Color]("green")
		Blue   = enum.M[Color]("blue")
		Colors = enum.New(Red, Green, Blue)
	)
	parsed := Colors.Parse("red")
	is.Equal(parsed, Red)
}
