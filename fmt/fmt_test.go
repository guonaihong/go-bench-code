package fmt2

import (
	"fmt"
	"testing"
)

func Benchmark_println(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("%s, %s\n", "hello", "world")
	}
}

func Benchmark_printf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("hello", "world")
	}
}
