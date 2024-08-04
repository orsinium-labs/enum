package enum_test

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/hcarriz/enum"
	"github.com/matryer/is"
)

type Color enum.Member[string]

// Ensure that Color is able to implement the driver.Valuer and sql.Scanner interface.
var _ driver.Valuer = (*Color)(nil)
var _ sql.Scanner = (*Color)(nil)

// Scan implements sql.Scanner.
func (c *Color) Scan(src any) error {
	var source string
	switch v := src.(type) {
	case string:
		source = v
	default:
		return fmt.Errorf("invalid type: %T", src)
	}
	*c = *Colors.Parse(source)
	return nil
}

// Value implements driver.Valuer.
func (c *Color) Value() (driver.Value, error) {
	return driver.Value(c.Property), nil
}

var (
	Red    = Color{"red"}
	Green  = Color{"green"}
	Blue   = Color{"blue"}
	Colors = enum.New(Red, Green, Blue)
)

func TestMember_Value(t *testing.T) {
	is := is.New(t)
	is.Equal(Red.Property, "red")
	is.Equal(Green.Property, "green")
	is.Equal(Blue.Property, "blue")
	is.Equal(enum.Member[string]{"blue"}.Property, "blue")
	is.Equal(enum.Member[int]{14}.Property, 14)
}

func TestEnum_Parse(t *testing.T) {
	is := is.New(t)
	var parsed *Color
	parsed = Colors.Parse("red")
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

func TestEnum_Choice(t *testing.T) {
	is := is.New(t)
	// Select a random color
	m := Colors.Choice(0)
	is.True(m != nil)
	is.True(Colors.Contains(*m))
	// Select a specific color using a specific random seed
	m = Colors.Choice(254)
	is.True(m != nil)
	is.Equal(*m, Red)
	// Select a specific color using a specific random seed
	m = Colors.Choice(1337)
	is.True(m != nil)
	is.Equal(*m, Green)
	// Select a specific color using a specific random seed
	m = Colors.Choice(42)
	is.True(m != nil)
	is.Equal(*m, Blue)
	// Selecting a random member from an empty Enum returns nil
	emptyEnums := enum.New[string, Color]()
	is.True(emptyEnums.Choice(0) == nil)
}

func TestEnum_Values(t *testing.T) {
	is := is.New(t)
	exp := []string{"red", "green", "blue"}
	is.Equal(Colors.Properties(), exp)
}

func TestEnum_Value(t *testing.T) {
	is := is.New(t)
	is.Equal(Colors.Property(Red), "red")
}

func TestEnum_Index(t *testing.T) {
	is := is.New(t)
	is.Equal(Colors.Index(Red), 0)
	is.Equal(Colors.Index(Green), 1)
	is.Equal(Colors.Index(Blue), 2)
}

func TestEnum_Index_Panic(t *testing.T) {
	is := is.New(t)
	defer func() {
		r := recover()
		is.Equal(r, "the given Member does not belong to this Enum")
	}()
	Colors.Index(Color{"purple"})
}

func TestBuilder(t *testing.T) {
	is := is.New(t)
	type Country enum.Member[string]
	var (
		b         = enum.NewBuilder[string, Country]()
		NL        = b.Add(Country{"Netherlands"})
		FR        = b.Add(Country{"France"})
		BE        = b.Add(Country{"Belgium"})
		Countries = b.Enum()
	)
	is.Equal(Countries.Members(), []Country{NL, FR, BE})
}

type BookValue struct {
	Title string
	ISBN  string
}

type Book enum.Member[BookValue]

var (
	EfficientGo     = Book{BookValue{"Efficient Go", "978-1098105716"}}
	ConcurrencyInGo = Book{BookValue{"Concurrency in Go", "978-1491941195"}}
	Books           = enum.New(EfficientGo, ConcurrencyInGo)
)

func (b BookValue) Equal(v BookValue) bool {
	return b.ISBN == v.ISBN
}

func TestParse(t *testing.T) {
	is := is.New(t)
	tests := []struct {
		isbn string
		want *Book
	}{
		{"978-1098105716", &EfficientGo},
		{"978-1491941195", &ConcurrencyInGo},
		{"invalid-isbn", nil},
	}
	for _, tt := range tests {
		t.Run(tt.isbn, func(t *testing.T) {
			v := BookValue{ISBN: tt.isbn}
			got := enum.Parse(Books, v)
			is.Equal(got, tt.want)
		})
	}
}
