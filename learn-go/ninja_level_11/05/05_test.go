package main

import "testing"

func TestRecursiveFactorial(t *testing.T) {
	x := recursiveFactorial(5)
	if x != 120 {
		t.Error("Expected:", 120, "Got:", x)
	}
}
func TestIterativeFactorial(t *testing.T) {
	x := iterativeFactorial(5)
	if x != 120 {
		t.Error("Expected:", 120, "Got:", x)
	}
}

func BenchmarkRecursiceFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recursiveFactorial(10)
	}
}

func BenchmarkIterativeFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iterativeFactorial(10)
	}
}
