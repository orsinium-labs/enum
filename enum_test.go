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
		Red   Color = "red"
		Green Color = "green"
	)
	var Colors = enum.New[string](Red, Green)
	parsed, ok := Colors.Parse("red")
	is.True(ok)
	is.Equal(parsed, Red)
}
