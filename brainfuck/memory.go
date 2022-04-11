package brainfuck

type memory struct {
	cells []byte
	ptr   int
}

func newMemory() *memory {
	return &memory{ptr: 0, cells: make([]byte, 1)}
}
