package brainfuck

func Interpret(code string) {
	commands := compile(code)
	exec(commands, newBuffer())
}

type command interface {
	execute(memory *buffer)
}

type commandStorage struct {
	cmds []command
	loop [][]command
}

func (c *commandStorage) add(cmd command) {
	depth := len(c.loop) - 1
	if depth >= 0 {
		c.loop[depth] = append(c.loop[depth], cmd)
	} else {
		c.cmds = append(c.cmds, cmd)
	}
}

func compile(code string) []command {
	var storage = &commandStorage{}
	for _, char := range code {
		addCommand(char, storage)
	}
	return storage.cmds
}

func exec(cmds []command, b *buffer) {
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
		c.loop = append(c.loop, []command{})
	case ']':
		cmd := loop{c.loop[len(c.loop)-1]}
		c.loop = c.loop[:len(c.loop)-1]
		c.add(cmd)
	case ' ', '\u000a':
		// ignore space and new line
	}
}
