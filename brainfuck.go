// package brainfuck defines interpretter for brinfuck programs
package brainfuck

import (
	"fmt"
	"strings"
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
type BFInterpretter struct {
	dataPointer  int
	inputValue   int
	inputPointer int
	outputValue  string

	tape      []int
	stackLoop []int
	inputTape []string
	output    []string
}

const maxTapeSize = 30000

func NewInterpretter() *BFInterpretter {
	interpretter := new(BFInterpretter)
	interpretter.tape = make([]int, 1, maxTapeSize)
	interpretter.stackLoop = make([]int, 1)
	interpretter.inputTape = make([]string, 1)
	return interpretter
}

// Rest https://stackoverflow.com/questions/30614165
func (bf *BFInterpretter) Reset() {
	if len(bf.tape) == 0 {
		return
	}

	bf.tape[0] = 0
	for bp := 1; bp < len(bf.tape); bp *= 2 {
		copy(bf.tape[bp:], bf.tape[:bp])
	}
	bf.dataPointer = 0

	bf.inputTape[0] = ""
	for bp := 1; bp < len(bf.inputTape); bp *= 2 {
		copy(bf.inputTape[bp:], bf.inputTape[:bp])
	}
	bf.inputPointer = 0
}

func (bf *BFInterpretter) Debug() (pointer int, tapeValue int, tapeLenght int) {
	return bf.dataPointer, bf.tape[bf.dataPointer], len(bf.tape)
}

func (bf *BFInterpretter) GetOutput() string {
	outputString := strings.Join(bf.output, "")
	return outputString
}

func (bf *BFInterpretter) Interpret(str string) string {
	// Working on the interpretter
	for _, inputByte := range str {
		bf.inputTape = append(bf.inputTape, string(inputByte))
	}

	for ; bf.inputPointer < len(bf.inputTape); bf.inputPointer++ {
		r := bf.inputTape[bf.inputPointer]
		switch r {
		case ".":
			fmt.Printf("%s", string(bf.tape[bf.dataPointer]))
			bf.output = append(bf.output, fmt.Sprintf("%s", string(bf.tape[bf.dataPointer])))
			bf.outputValue = fmt.Sprintf("%+q", bf.tape[bf.dataPointer])
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
		case "[":
			bf.stackLoop = append(bf.stackLoop, bf.inputPointer)
			break
		case "]":
			if tapeValue := bf.tape[bf.dataPointer]; tapeValue == 0 {
				lenStackLoop := len(bf.stackLoop) - 1
				// Delete last element, ... is important
				// s = append(s[:0], s[:len(s) - 1]...)
				bf.stackLoop = append(bf.stackLoop[:0], bf.stackLoop[:lenStackLoop]...)
			} else {
				// Reset inputPointer to beginning of the loop
				bf.inputPointer = bf.stackLoop[len(bf.stackLoop)-1]
			}
		}
	}
	return ""
}
