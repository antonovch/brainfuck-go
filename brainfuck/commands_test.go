package brainfuck

import (
	"reflect"
	"testing"
)

func TestLoop_execute(t *testing.T) {
	type fields struct {
		cmds []command
	}
	type args struct {
		memory *memory
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected memory
	}{
		{
			"skips executing commands if memory current cell value is zero",
			fields{[]command{forward{}, increment{}}},
			args{memory: newMemory()},
			*newMemory(),
		},
		{
			"executes all commands if memory current cell value is non-zero",
			fields{[]command{
				forward{}, increment{}, increment{}, backward{}, decrement{},
			}},
			args{memory: &memory{[]byte{1}, 0}},
			memory{[]byte{0, 2}, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := loop{
				cmds: tt.fields.cmds,
			}
			l.execute(tt.args.memory)
			if !reflect.DeepEqual(*tt.args.memory, tt.expected) {
				t.Errorf("Expected %v \n but have %v", tt.expected, *tt.args.memory)
			}
		})
	}
}
