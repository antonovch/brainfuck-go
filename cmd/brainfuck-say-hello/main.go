package main

import (
	"fmt"
	"github.com/antonovch/brainfuck-go/brainfuck"
	"io/ioutil"
	"os"
)

func main() {
	var code string

	if len(os.Args) != 3 {
		fmt.Println("Wrong number of input arguments.")
		return
	}

	switch os.Args[1] {
	case "-f":
		content, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		code = string(content)
	case "-i":
		code = os.Args[2]
	default:
		fmt.Println("Allowed options are -f [filename] or -i [code string].")
	}

	fmt.Println(brainfuck.Interpret(code))
}
