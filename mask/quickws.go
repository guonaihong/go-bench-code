package mask

import (
	"encoding/binary"
	"reflect"
	"unsafe"
)

func maskSlow2(payload []byte, key uint32) {
	var maskVal [4]byte
	binary.LittleEndian.PutUint32(maskVal[:], key)
	for i := 0; i < len(payload); i++ {
		payload[i] ^= maskVal[i%4]
	}
}

//go:nosplit
func maskQuickws(payload []byte, key uint32) {
	if len(payload) == 0 {
		return
	}

	base := (*reflect.SliceHeader)(unsafe.Pointer(&payload)).Data
	last := base + uintptr(len(payload))
	if len(payload) >= 8 {
		key64 := uint64(key)<<32 | uint64(key)

		var v *uint64
		for base+128 <= last {

			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
		}

		if base == last {
			return
		}

		if base+64 <= last {

			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
		}

		if base+32 <= last {
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
		}

		if base+16 <= last {
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
		}

		if base+8 <= last {
			v = (*uint64)(unsafe.Pointer(base))
			*v ^= key64
			base += 8
		}

		if base == last {
			return
		}
	}

	if base+4 <= last {
		v := (*uint32)(unsafe.Pointer(base))
		*v ^= key
		base += 4
	}

	if base == last {
		return
	}
	if base < last {
		maskSlow2(payload[len(payload)-int(last-base):], key)
	}
}
