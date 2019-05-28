package main

import "testing"

func Benchmark_make_copy(b *testing.B) {
	testData := []string{"hello", "world", "123"}
	for i := 0; i < b.N; i++ {
		t2 := make([]string, len(testData))
		copy(t2, testData)
	}
}

func Benchmark_append(b *testing.B) {
	testData := []string{"hello", "world", "123"}
	for i := 0; i < b.N; i++ {
		_ = append([]string{}, testData...)
	}
}

func Benchmark_make_append(b *testing.B) {
	testData := []string{"hello", "world", "123"}
	for i := 0; i < b.N; i++ {
		_ = append(make([]string, 0, len(testData)), testData...)
	}
}
