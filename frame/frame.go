package frame

import (
	"encoding/binary"
	"errors"
	"io"
)

const (
	// 根据5.2描述, 满打满算, 最多14字节
	maxFrameHeaderSize = 14
)

var ErrFramePayloadLength = errors.New("error frame payload length")

type frameHeader struct {
	payloadLen int64
	opcode     Opcode
	maskValue  [4]byte
	rsv1       bool
	rsv2       bool
	rsv3       bool
	mask       bool
	fin        bool
}

func readHeader(r io.Reader) (h frameHeader, size int, err error) {
	var headArray [maxFrameHeaderSize]byte
	head := headArray[:2]

	n, err := io.ReadFull(r, head)
	if err != nil {
		return
	}
	if n != 2 {
		err = io.ErrUnexpectedEOF
		return
	}
	size = 2

	h.fin = head[0]&(1<<7) > 0
	h.rsv1 = head[0]&(1<<6) > 0
	h.rsv2 = head[0]&(1<<5) > 0
	h.rsv3 = head[0]&(1<<4) > 0
	h.opcode = Opcode(head[0] & 0xF)

	have := 0
	h.mask = head[1]&(1<<7) > 0
	if h.mask {
		have += 4
		size += 4
	}

	h.payloadLen = int64(head[1] & 0x7F)

	switch {
	// 长度
	case h.payloadLen >= 0 && h.payloadLen <= 125:
		if h.payloadLen == 0 && !h.mask {
			return
		}
	case h.payloadLen == 126:
		// 2字节长度
		have += 2
		size += 2
	case h.payloadLen == 127:
		// 8字节长度
		have += 8
		size += 8
	default:
		// 预期之外的, 直接报错
		return h, 0, ErrFramePayloadLength
	}

	head = head[:have]
	_, err = io.ReadFull(r, head)
	if err != nil {
		return
	}

	switch h.payloadLen {
	case 126:
		h.payloadLen = int64(binary.BigEndian.Uint16(head[:2]))
		head = head[2:]
	case 127:
		h.payloadLen = int64(binary.BigEndian.Uint64(head[:8]))
		head = head[8:]
	}

	if h.mask {
		copy(h.maskValue[:], head)
	}

	return
}

func readHeader2(r io.Reader, headArray [maxFrameHeaderSize]byte) (h frameHeader, size int, err error) {
	// var headArray [maxFrameHeaderSize]byte
	head := headArray[:2]

	n, err := io.ReadFull(r, head)
	if err != nil {
		return
	}
	if n != 2 {
		err = io.ErrUnexpectedEOF
		return
	}
	size = 2

	h.fin = head[0]&(1<<7) > 0
	h.rsv1 = head[0]&(1<<6) > 0
	h.rsv2 = head[0]&(1<<5) > 0
	h.rsv3 = head[0]&(1<<4) > 0
	h.opcode = Opcode(head[0] & 0xF)

	have := 0
	h.mask = head[1]&(1<<7) > 0
	if h.mask {
		have += 4
		size += 4
	}

	h.payloadLen = int64(head[1] & 0x7F)

	switch {
	// 长度
	case h.payloadLen >= 0 && h.payloadLen <= 125:
		if h.payloadLen == 0 && !h.mask {
			return
		}
	case h.payloadLen == 126:
		// 2字节长度
		have += 2
		size += 2
	case h.payloadLen == 127:
		// 8字节长度
		have += 8
		size += 8
	default:
		// 预期之外的, 直接报错
		return h, 0, ErrFramePayloadLength
	}

	head = head[:have]
	_, err = io.ReadFull(r, head)
	if err != nil {
		return
	}

	switch h.payloadLen {
	case 126:
		h.payloadLen = int64(binary.BigEndian.Uint16(head[:2]))
		head = head[2:]
	case 127:
		h.payloadLen = int64(binary.BigEndian.Uint64(head[:8]))
		head = head[8:]
	}

	if h.mask {
		copy(h.maskValue[:], head)
	}

	return
}

func readHeader3(r io.Reader, headArray *[maxFrameHeaderSize]byte) (h frameHeader, size int, err error) {
	// var headArray [maxFrameHeaderSize]byte
	head := (*headArray)[:2]

	n, err := io.ReadFull(r, head)
	if err != nil {
		return
	}
	if n != 2 {
		err = io.ErrUnexpectedEOF
		return
	}
	size = 2

	h.fin = head[0]&(1<<7) > 0
	h.rsv1 = head[0]&(1<<6) > 0
	h.rsv2 = head[0]&(1<<5) > 0
	h.rsv3 = head[0]&(1<<4) > 0
	h.opcode = Opcode(head[0] & 0xF)

	have := 0
	h.mask = head[1]&(1<<7) > 0
	if h.mask {
		have += 4
		size += 4
	}

	h.payloadLen = int64(head[1] & 0x7F)

	switch {
	// 长度
	case h.payloadLen >= 0 && h.payloadLen <= 125:
		if h.payloadLen == 0 && !h.mask {
			return
		}
	case h.payloadLen == 126:
		// 2字节长度
		have += 2
		size += 2
	case h.payloadLen == 127:
		// 8字节长度
		have += 8
		size += 8
	default:
		// 预期之外的, 直接报错
		return h, 0, ErrFramePayloadLength
	}

	head = head[:have]
	_, err = io.ReadFull(r, head)
	if err != nil {
		return
	}

	switch h.payloadLen {
	case 126:
		h.payloadLen = int64(binary.BigEndian.Uint16(head[:2]))
		head = head[2:]
	case 127:
		h.payloadLen = int64(binary.BigEndian.Uint64(head[:8]))
		head = head[8:]
	}

	if h.mask {
		copy(h.maskValue[:], head)
	}

	return
}

func readHeader4(r io.Reader, headArray *[maxFrameHeaderSize]byte, h *frameHeader) (size int, err error) {
	// var headArray [maxFrameHeaderSize]byte
	head := (*headArray)[:2]

	n, err := io.ReadFull(r, head)
	if err != nil {
		return
	}
	if n != 2 {
		err = io.ErrUnexpectedEOF
		return
	}
	size = 2

	h.fin = head[0]&(1<<7) > 0
	h.rsv1 = head[0]&(1<<6) > 0
	h.rsv2 = head[0]&(1<<5) > 0
	h.rsv3 = head[0]&(1<<4) > 0
	h.opcode = Opcode(head[0] & 0xF)

	have := 0
	h.mask = head[1]&(1<<7) > 0
	if h.mask {
		have += 4
		size += 4
	}

	h.payloadLen = int64(head[1] & 0x7F)

	switch {
	// 长度
	case h.payloadLen >= 0 && h.payloadLen <= 125:
		if h.payloadLen == 0 && !h.mask {
			return
		}
	case h.payloadLen == 126:
		// 2字节长度
		have += 2
		size += 2
	case h.payloadLen == 127:
		// 8字节长度
		have += 8
		size += 8
	default:
		// 预期之外的, 直接报错
		return 0, ErrFramePayloadLength
	}

	head = head[:have]
	_, err = io.ReadFull(r, head)
	if err != nil {
		return
	}

	switch h.payloadLen {
	case 126:
		h.payloadLen = int64(binary.BigEndian.Uint16(head[:2]))
		head = head[2:]
	case 127:
		h.payloadLen = int64(binary.BigEndian.Uint64(head[:8]))
		head = head[8:]
	}

	if h.mask {
		copy(h.maskValue[:], head)
	}

	return
}
