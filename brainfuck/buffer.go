package brainfuck

type buffer struct {
	cells []byte
	ptr   int
}

func newBuffer() *buffer {
	return &buffer{ptr: 0, cells: make([]byte, 1)}
}
