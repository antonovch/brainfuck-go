package main

import "github.com/antonovch/brainfuck-go/brainfuck"

func main() {
	helloWorldCode := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
	brainfuck.Interpret(helloWorldCode)
}
