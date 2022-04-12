package brainfuck

/*
	forward defines an execute method that moves memory's pointer one step forward ('>' operator).
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
	backward defines an execute method that moves memory's pointer one step backward ('<' operator).
*/
type backward struct{}

func (backward) execute(m *memory) {
	m.ptr--
}

/*
	increment defines an execute method that increments memory's value pointed to ('+' operator).
*/
type increment struct{}

func (increment) execute(m *memory) {
	m.cells[m.ptr]++
}

/*
	decrement defines an execute method that decrements memory's value pointed to ('-' operator).
*/
type decrement struct{}

func (decrement) execute(m *memory) {
	m.cells[m.ptr]--
}

/*
	output defines an execute method that prints memory's value pointed to as an ASCII char ('.' operator).
*/
type output struct{}

func (output) execute(m *memory) {
	m.output += string(m.cells[m.ptr])
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
