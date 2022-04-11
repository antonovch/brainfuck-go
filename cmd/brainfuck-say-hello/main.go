package main

import (
	"fmt"
	"github.com/antonovch/brainfuck-go/brainfuck"
)

func main() {
	helloWorldCode := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
	fmt.Println(brainfuck.Interpret(helloWorldCode))
}
