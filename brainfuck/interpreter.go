package brainfuck

/*
	Interpret interprets brainfuck code string and prints
	the output to the console through brainfuck's '.' directive.
	The function does two things:

		1. transforms the code string into a slice of commands;

		2. runs the commands in order.
*/
func Interpret(code string) {
	commands := compile(code)
	exec(commands, newMemory())
}

type command interface {
	execute(memory *memory)
}

type commandStorage struct {
	cmds []command
	loop [][]command
}

/*
	add adds new command to commandStorage. If loop slice of slices is not
	empty (nested loops), appends the command to the last cell, otherwise
	cmd command is appended directly to cmds. loop's subslices are
	moved to type loop structure and finally appended to cmds elsewhere.
*/
func (c *commandStorage) add(cmd command) {
	depth := len(c.loop) - 1
	if depth >= 0 {
		c.loop[depth] = append(c.loop[depth], cmd)
	} else {
		c.cmds = append(c.cmds, cmd)
	}
}

/*
	compile transforms the code string into a slice of commands.
*/
func compile(code string) []command {
	var storage = &commandStorage{}
	for _, char := range code {
		addCommand(char, storage)
	}
	return storage.cmds
}

/*
	exec executes a given slice of commands in sequence.
*/
func exec(cmds []command, m *memory) {
	for _, cmd := range cmds {
		cmd.execute(m)
	}
}

/*
	addCommand adds new commands to commandStorage with
	the proper logic.
*/
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
	}
}
