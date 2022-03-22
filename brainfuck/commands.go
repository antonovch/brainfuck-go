package brainfuck

import "fmt"

type forward struct{}

func (forward) execute(b *buffer) {
	if len(b.cells)-1 == b.ptr {
		b.cells = append(b.cells, make([]byte, len(b.cells))...)
	}
	b.ptr++
}

type backward struct{}

func (backward) execute(b *buffer) {
	b.ptr--
}

type increment struct{}

func (increment) execute(b *buffer) {
	b.cells[b.ptr]++
}

type decrement struct{}

func (decrement) execute(b *buffer) {
	b.cells[b.ptr]--
}

type output struct{}

func (output) execute(b *buffer) {
	fmt.Print(string(b.cells[b.ptr]))
}

type loop struct {
	cmds []multi
}

func (l loop) execute(b *buffer) {
	for b.cells[b.ptr] != 0 {
		for _, c := range l.cmds {
			c.execute(b)
		}
	}
}

type multi struct {
	cmd   command
	count int
}

func (m multi) execute(b *buffer) {
	for ; m.count > 0; m.count-- {
		m.cmd.execute(b)
	}
}

func newMulti(cmd command) multi {
	return multi{cmd, 1}
}
