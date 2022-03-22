package brainfuck

import "fmt"

/*
	Move buffer's pointer one step forward ('>' operator).
*/
type forward struct{}

func (forward) execute(b *buffer) {
	// when the end of the buffer is reached, it'll grow to twice its size
	if len(b.cells)-1 == b.ptr {
		b.cells = append(b.cells, make([]byte, len(b.cells))...)
	}
	b.ptr++
}

/*
	Move buffer's pointer one step backward ('<' operator).
*/
type backward struct{}

func (backward) execute(b *buffer) {
	b.ptr--
}

/*
	Increment buffer's value pointed to ('+' operator).
*/
type increment struct{}

func (increment) execute(b *buffer) {
	b.cells[b.ptr]++
}

/*
	Decrement buffer's value pointed to ('-' operator).
*/
type decrement struct{}

func (decrement) execute(b *buffer) {
	b.cells[b.ptr]--
}

/*
	Print buffer's value pointed to as an ASCII char ('.' operator).
*/
type output struct{}

func (output) execute(b *buffer) {
	fmt.Print(string(b.cells[b.ptr]))
}

/*
	loop struct contains slice of commands, specified in between
	'[' and ']' operators. They are executed sequentially in a loop.
*/
type loop struct {
	cmds []command
}

func (l loop) execute(b *buffer) {
	for b.cells[b.ptr] != 0 {
		for _, c := range l.cmds {
			c.execute(b)
		}
	}
}
