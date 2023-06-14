package frame

type Opcode uint8

const (
	Continuation Opcode = iota
	Text
	Binary
	// 3 - 7保留
	_ // 3
	_
	_ // 5
	_
	_ // 7
	Close
	Ping
	Pong
)
