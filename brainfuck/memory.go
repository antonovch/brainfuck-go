package brainfuck

type memory struct {
	cells  []byte
	output string
	ptr    int
}

func newMemory() *memory {
	return &memory{
		ptr:   0,
		cells: make([]byte, 1),
	}
}
