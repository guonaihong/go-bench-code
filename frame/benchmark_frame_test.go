package frame

import (
	"bytes"
	"testing"
)

var (
	noMaskData   = []byte{0x81, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f}
	haveMaskData = []byte{0x81, 0x85, 0x37, 0xfa, 0x21, 0x3d, 0x7f, 0x9f, 0x4d, 0x51, 0x58}
)

func Benchmark_ReadFrame(b *testing.B) {
	r := bytes.NewReader(noMaskData)
	for i := 0; i < b.N; i++ {

		r.Reset(noMaskData)
		_, _, err := readHeader(r)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ReadFrame2(b *testing.B) {
	r := bytes.NewReader(noMaskData)
	var headArray [maxFrameHeaderSize]byte
	for i := 0; i < b.N; i++ {

		r.Reset(noMaskData)
		_, _, err := readHeader2(r, headArray)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ReadFrame3(b *testing.B) {
	r := bytes.NewReader(noMaskData)
	var headArray [maxFrameHeaderSize]byte
	for i := 0; i < b.N; i++ {

		r.Reset(noMaskData)
		_, _, err := readHeader3(r, &headArray)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ReadFrame4(b *testing.B) {
	r := bytes.NewReader(noMaskData)
	var headArray [maxFrameHeaderSize]byte
	var f frameHeader
	for i := 0; i < b.N; i++ {

		r.Reset(noMaskData)
		_, err := readHeader4(r, &headArray, &f)
		if err != nil {
			b.Fatal(err)
		}
	}
}
