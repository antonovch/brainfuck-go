/*
Package brainfuck provides functionality to interpret any classic brainfuck
program.

Functions:
	Interpret
		- takes a string of brainfuck code that consists of symbols +-<>.[]
		- returns a string with the code's output

Example usage:
	brainfuck.Interpret("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
*/
package brainfuck
