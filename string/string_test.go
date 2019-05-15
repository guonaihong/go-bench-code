package string

import (
	"bytes"
	"strings"
	"testing"
)

func Benchmark_Add(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += "hello"
	}
}

func Benchmark_BytesWrite(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		buf.WriteString("hello")
	}
}

func Benchmark_builder(b *testing.B) {
	var s strings.Builder

	for i := 0; i < b.N; i++ {
		s.WriteString("hello")
	}
}

func Benchmark_BytesWriteMem(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		buf.WriteString("hello")
		buf.String()
	}
}

func Benchmark_builderMem(b *testing.B) {
	var s strings.Builder

	for i := 0; i < b.N; i++ {
		s.WriteString("hello")
		s.String()
	}
}
