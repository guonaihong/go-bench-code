package copytest

import "testing"

func Benchmark_Copy1024_1(b *testing.B) {
	buf := make([]byte, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(buf, buf)
	}
}

func Benchmark_Copy1024_2(b *testing.B) {
	buf := make([]byte, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(buf, buf[2:])
	}
}
