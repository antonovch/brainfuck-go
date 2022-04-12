# brainfuck-go
A simple realisation of [`Brainfuck`](https://en.wikipedia.org/wiki/Brainfuck) language interpreter in Go Language.
The program maintains a byte array, whose values can be manipulated using the _supported operations_.

### Supported operations

1. `>`, increment the data pointer,
2. `<`, decrement thee data pointer,
3. `+`, increment the byte at the data pointer,
4. `-`, decrement the byte at the data pointer,
5. `.`, output the byte at the data pointer,
6. `[`, start loop: if the byte at the data pointer is zero, then jump forward to the command after the matching `]`, 
   else move the pointer forward,
7. `]`, end loop: if the byte at the data pointer is nonzero, then jump back to the command after the matching `[`,
   else move the pointer forward.

## Installation
### Required software
* [`go1.17`](https://go.dev/dl/) or newer

To download and install the package use `go get` command:
```bash
go get -u github.com/antonovch/brainfuck-go
```
To install already downloaded module, use `go install` (when working from the repo root):
```bash
go install ./brainfuck
```

## Tests
Test can be run the usual way (from the package folder):
```bash
$ go test -cover
Hello World!
PASS
coverage: 100.0% of statements
ok      github.com/antonovch/brainfuck-go/brainfuck     0.729s
```

## Example usage
The file `cmd/brainfuck-say-hello/main.go` provides a program code that can be installed using
```bash
go install ./cmd/brainfuck-say-hello 
```
and invoked from the command line either with a filename and a `-f` flag:
```bash
brainfuck-say-hello -f "/path/to/file.b"
```
or a string and a `-i` flag:
```bash
brainfuck-say-hello -i "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.-
-------.>>+.>++."
```