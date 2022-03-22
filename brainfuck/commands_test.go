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
		buffer *buffer
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected buffer
	}{
		{
			"skips executing commands if buffer current cell value is zero",
			fields{[]command{forward{}, increment{}}},
			args{buffer: newBuffer()},
			*newBuffer(),
		},
		{
			"executes all commands if buffer current cell value is non-zero",
			fields{[]command{
				forward{}, increment{}, increment{}, backward{}, decrement{},
			}},
			args{buffer: &buffer{[]byte{1}, 0}},
			buffer{[]byte{0, 2}, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := loop{
				cmds: tt.fields.cmds,
			}
			l.execute(tt.args.buffer)
			if !reflect.DeepEqual(*tt.args.buffer, tt.expected) {
				t.Errorf("Expected %v \n but have %v", tt.expected, *tt.args.buffer)
			}
		})
	}
}
