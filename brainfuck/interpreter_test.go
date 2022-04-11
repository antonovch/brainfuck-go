package brainfuck

import (
	"reflect"
	"testing"
)

func TestInterpret(t *testing.T) {
	type args struct {
		program string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{"hello world",
			args{"++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."},
			"Hello World!\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Interpret(tt.args.program)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v \n but have %v", tt.expected, result)
			}
		})
	}
}
