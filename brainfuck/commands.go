package brainfuck

import "fmt"

/*
	Move memory's pointer one step forward ('>' operator).
*/
type forward struct{}

func (forward) execute(m *memory) {
	// when the end of the memory is reached, it'll grow to twice its size
	if len(m.cells)-1 == m.ptr {
		m.cells = append(m.cells, make([]byte, len(m.cells))...)
	}
	m.ptr++
}

/*
	Move memory's pointer one step backward ('<' operator).
*/
type backward struct{}

func (backward) execute(m *memory) {
	m.ptr--
}

/*
	Increment memory's value pointed to ('+' operator).
*/
type increment struct{}

func (increment) execute(m *memory) {
	m.cells[m.ptr]++
}

/*
	Decrement memory's value pointed to ('-' operator).
*/
type decrement struct{}

func (decrement) execute(m *memory) {
	m.cells[m.ptr]--
}

/*
	Print memory's value pointed to as an ASCII char ('.' operator).
*/
type output struct{}

func (output) execute(m *memory) {
	fmt.Print(string(m.cells[m.ptr]))
}

/*
	loop struct contains slice of commands, specified in between
	'[' and ']' operators. They are executed sequentially in a loop.
*/
type loop struct {
	cmds []command
}

func (l loop) execute(m *memory) {
	for m.cells[m.ptr] != 0 {
		for _, c := range l.cmds {
			c.execute(m)
		}
	}
}
