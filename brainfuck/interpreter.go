package brainfuck

func Interpret(code string) {
	commands := compile(code)
	exec(commands, newBuffer())
}

type command interface {
	execute(memory *buffer)
}

type commandStorage struct {
	cmds []multi
	loop [][]multi
}

func (c *commandStorage) add(cmd command) {
	depth := len(c.loop) - 1
	if depth >= 0 {
		c.loop[depth] = multiAppend(c.loop[depth], cmd)
	} else {
		c.cmds = multiAppend(c.cmds, cmd)
	}
}

func multiAppend(cmds []multi, c command) []multi {
	if len(cmds) > 0 && cmds[len(cmds)-1].cmd == c {
		cmds[len(cmds)-1].count++
	} else {
		cmds = append(cmds, newMulti(c))
	}
	return cmds
}

func compile(code string) []multi {
	var storage = &commandStorage{}
	for _, char := range code {
		addCommand(char, storage)
	}
	return storage.cmds
}

func exec(cmds []multi, b *buffer) {
	for _, cmd := range cmds {
		cmd.execute(b)
	}
}

func addCommand(char rune, c *commandStorage) {
	switch char {
	case '>':
		c.add(forward{})
	case '<':
		c.add(backward{})
	case '+':
		c.add(increment{})
	case '-':
		c.add(decrement{})
	case '.':
		c.add(output{})
	case '[':
		c.loop = append(c.loop, []multi{})
	case ']':
		cmd := loop{c.loop[len(c.loop)-1]}
		c.loop = c.loop[:len(c.loop)-1]
		c.add(cmd)
	}
}
