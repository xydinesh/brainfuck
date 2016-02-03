// package brainfuck defines interpretter for brinfuck programs
package brainfuck

import (
	"fmt"
)

/*
Any character not "><+-.,[]" (excluding quotation marks) is ignored.

Brainfuck is represented by an array with 30,000 cells initialized to zero
and a data pointer pointing at the current cell.

There are eight commands:
+ : Increments the value at the current cell by one.
- : Decrements the value at the current cell by one.
> : Moves the data pointer to the next cell (cell on the right).
< : Moves the data pointer to the previous cell (cell on the left).
. : Prints the ASCII value at the current cell (i.e. 65 = 'A').
, : Reads a single input character into the current cell.
[ : If the value at the current cell is zero, skips to the corresponding ] .
    Otherwise, move to the next instruction.
] : If the value at the current cell is zero, move to the next instruction.
    Otherwise, move backwards in the instructions to the corresponding [ .

[ and ] form a while loop. Obviously, they must be balanced.
*/
type BFInterpreter struct {
	dataPointer int
	tape        []int
	inputValue  int
}

const maxTapeSize = 30000

func NewInterpretter() *BFInterpreter {
	interpretter := new(BFInterpreter)
	interpretter.tape = make([]int, 1, maxTapeSize)
	return interpretter
}

// Rest https://stackoverflow.com/questions/30614165
func (bf *BFInterpreter) Reset() {
	if len(bf.tape) == 0 {
		return
	}

	bf.tape[0] = 0
	for bp := 1; bp < len(bf.tape); bp *= 2 {
		copy(bf.tape[bp:], bf.tape[:bp])
	}
	bf.dataPointer = 0
}

func (bf *BFInterpreter) Interpret(str string) string {
	// Working on the interpretter
	for _, r := range str {
		switch string(r) {
		case ".":
			return fmt.Sprintf("%+q", bf.tape[bf.dataPointer])
		case "+":
			bf.tape[bf.dataPointer]++
			break
		case "-":
			bf.tape[bf.dataPointer]--
			break
		case ">":
			if bf.dataPointer++; bf.dataPointer >= len(bf.tape) {
				bf.tape = append(bf.tape, 0)
			}
			break
		case "<":
			if bf.dataPointer > 0 {
				bf.dataPointer--
			}
			break
		case ",":
			fmt.Scanf("%d\n", bf.inputValue)
			bf.tape[bf.dataPointer] = bf.inputValue
			break
		}
	}
	return str
}
