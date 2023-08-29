package enum_test

import (
	"fmt"

	"github.com/orsinium-labs/enum"
)

func ExampleNew() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	fmt.Printf("%#v\n", Colors)
	// Output: enum.New[string](enum_test.Color{"red"}, enum_test.Color{"green"}, enum_test.Color{"blue"})
}

func ExampleEnum_String() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	fmt.Println(Colors)
	// Output: red, green, blue
}

func ExampleEnum_GoString() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	fmt.Printf("%#v\n", Colors)
	// Output: enum.New[string](enum_test.Color{"red"}, enum_test.Color{"green"}, enum_test.Color{"blue"})
}

func ExampleEnum_Parse() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	parsed := Colors.Parse("red")
	fmt.Printf("%#v\n", parsed)
	// Output: &enum_test.Color{Value:"red"}
}

func ExampleEnum_Contains() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	contains := Colors.Contains(Red)
	fmt.Println(contains)
	// Output: true
}

func ExampleEnum_Empty() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	empty := Colors.Empty()
	fmt.Println(empty)
	// Output: false
}

func ExampleEnum_Index() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	index := Colors.Index(Green)
	fmt.Println(index)
	// Output: 1
}

func ExampleEnum_Value() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	value := Colors.Value(Green)
	fmt.Println(value)
	// Output: green
}

func ExampleEnum_Len() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	length := Colors.Len()
	fmt.Println(length)
	// Output: 3
}

func ExampleEnum_Members() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	members := Colors.Members()
	fmt.Println(members)
	// Output: [{red} {green} {blue}]
}

func ExampleEnum_Values() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	values := Colors.Values()
	fmt.Println(values)
	// Output: [red green blue]
}

func ExampleEnum_TypeName() {
	type Color enum.Member[string]

	var (
		Red    = Color{"red"}
		Green  = Color{"green"}
		Blue   = Color{"blue"}
		Colors = enum.New[string](Red, Green, Blue)
	)

	tname := Colors.TypeName()
	fmt.Println(tname)
	// Output: string
}
