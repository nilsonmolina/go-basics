package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(3)
	}
}
func BenchmarkYearsAlt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsAlt(3)
	}
}

func ExampleYears() {
	fmt.Println(Years(3))
	// Output:
	// 21
}
func ExampleYearsAlt() {
	fmt.Println(YearsAlt(3))
	// Output:
	// 21
}

func TestYears(t *testing.T) {
	got := Years(10)
	want := 70

	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
func TestYearsAlt(t *testing.T) {
	got := YearsAlt(10)
	want := 70

	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
