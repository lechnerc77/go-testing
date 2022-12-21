package calculator_test

import (
	"testing"

	"github.com/lechnerc77/gotestdemo/calculator"
)

func benchmarkIsPrime(b *testing.B, value int) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		calculator.IsPrime(value)
	}
}

func BenchmarkIsPrime2(b *testing.B) {
	benchmarkIsPrime(b, 2)
}

func BenchmarkIsPrime101(b *testing.B) {
	benchmarkIsPrime(b, 101)
}

func BenchmarkIsPrime1235868985(b *testing.B) {
	benchmarkIsPrime(b, 1235868985)
}
