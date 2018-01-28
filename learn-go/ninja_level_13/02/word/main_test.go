package word

import (
	"fmt"
	"testing"

	"github.com/nilsonmolina/go-basics/learn-go/ninja_level_13/02/quote"
)

func TestCount(t *testing.T) {
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got:", v, "want:", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got:", v, "want:", 1)
			}
		case "three":
			if v != 3 {
				t.Error("got:", v, "want:", 3)
			}
		}
	}
}

func TestUseCount(t *testing.T) {
	got := Count("one two three")
	want := 3

	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
