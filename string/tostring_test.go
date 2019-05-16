package string

import (
	"fmt"
	"strconv"
	"testing"
)

func Benchmark_Sprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", i)
	}
}

func Benchmark_Itoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(i)
	}
}
