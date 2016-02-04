package brainfuck

import "testing"

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
func TestBFStruct(t *testing.T) {
	t.Log("Testing interpreter structure")
	bf := BFInterpretter{}
	if pointer := bf.dataPointer; pointer != 0 {
		t.Errorf("Expected dataPointer at 0, but it was %d instead", pointer)
	}

	if tape := bf.tape; tape != nil {
		t.Errorf("Expected tape slice to be nil")
	}

	if maxTapeSize != 30000 {
		t.Errorf("maxTapeSize expected to be 30,000 but it was %d", maxTapeSize)
	}
}

func TestNewInterpretter(t *testing.T) {
	t.Log("Test creating new interpertter")
	bf := NewInterpretter()
	if pointer := bf.dataPointer; pointer != 0 {
		t.Errorf("Expected dataPointer at 0, but it was %d instead", pointer)
	}

	if tape := bf.tape; tape == nil {
		t.Errorf("Expected tape slice to be nil")
	}

	if inputTape := bf.inputTape; inputTape == nil {
		t.Errorf("Expected inputTape slice to be nil")
	}

	if inputPointer := bf.inputPointer; inputPointer != 0 {
		t.Errorf("Expected inputPointer at 0, but it was %d instead", inputPointer)
	}

	if stack := bf.stackLoop; stack == nil {
		t.Errorf("Expected stackLoop slice to be nil")
	}

	if tapeCap := cap(bf.tape); tapeCap != maxTapeSize {
		t.Errorf("Expected tape capacity to be %d, but it was %d", maxTapeSize, tapeCap)
	}

	if tapeSize := len(bf.tape); tapeSize != 1 {
		t.Errorf("Expected tape capacity to be 1, but it was %d", tapeSize)
	}
}

func TestIneterpretPrint(t *testing.T) {
	t.Log("Test BF Interpretter print with .")
	bf := NewInterpretter()
	bf.Interpret(".")
	if output := bf.outputValue; output != "'\\x00'" {
		t.Fatalf("Expected to print 0, but it was %q", output)
	}
}

func TestIneterpretPlus(t *testing.T) {
	t.Log("Test BF Interpretter incremant with +")
	bf := NewInterpretter()
	bf.Interpret("+.")
	if output := bf.outputValue; output != "'\\x01'" {
		t.Errorf("Expected to print 1, but it was %q", output)
	}

	bf.Interpret("++ +.")
	if output := bf.outputValue; output != "'\\x04'" {
		t.Errorf("Expected to print 4, but it was %q", output)
	}
}

func TestInerpretRest(t *testing.T) {
	t.Log("Testing interpretter reset function")
	bf := NewInterpretter()
	bf.Interpret("+.")
	if output := bf.outputValue; output != "'\\x01'" {
		t.Errorf("Expected to print 1, but it was %q", output)
	}

	bf.Reset()
	bf.Interpret(".")
	if output := bf.outputValue; output != "'\\x00'" {
		t.Errorf("Expected to print 0, but it was %q", output)
	}
}

func TestInterpretMinus(t *testing.T) {
	t.Log("Testing BF Interpretter decrement with -")

	bf := NewInterpretter()
	bf.Interpret("+.")
	if output := bf.outputValue; output != "'\\x01'" {
		t.Errorf("Expected to print 1, but it was %q", output)
	}

	bf.Interpret("++")
	bf.Interpret("-.")
	if output := bf.outputValue; output != "'\\x02'" {
		t.Errorf("Expected to print 2, but it was %q", output)
	}
	bf.Interpret("+++ ++ --.")
	if output := bf.outputValue; output != "'\\x05'" {
		t.Errorf("Expected to print 5, but it was %q", output)
	}

}

func TestInterpretA(t *testing.T) {
	t.Log("Testing print A")
	bf := NewInterpretter()
	for i := 0; i < 6; i++ {
		bf.Interpret("+++++ +++++")
	}

	bf.Interpret("+++++ .")
	if output := bf.outputValue; output != "'A'" {
		t.Errorf("Expected to print A, but it was %+q", output)
	}
}

func TestInterpretForward(t *testing.T) {
	t.Log("Testing BF Interpretter data pointer moving forward")
	bf := NewInterpretter()
	bf.Interpret("+.")
	if output := bf.outputValue; output != "'\\x01'" {
		t.Errorf("Expected to print 1, but it was %q", output)
	}

	bf.Interpret("> +.")
	if output := bf.outputValue; output != "'\\x01'" {
		t.Errorf("Expected to print 1, but it was %q", output)
	}
}

func TestInterpretBackward(t *testing.T) {
	t.Log("Test BF Interpretter data pointer moving backward")
	bf := NewInterpretter()
	bf.Interpret("+.")
	if output := bf.outputValue; output != "'\\x01'" {
		t.Errorf("Expected to print 1, but it was %q", output)
	}

	bf.Interpret("> +++++.")
	if output := bf.outputValue; output != "'\\x05'" {
		t.Errorf("Expected to print 5, but it was %q", output)
	}

	bf.Interpret("<.")
	if output := bf.outputValue; output != "'\\x01'" {
		t.Errorf("Expected to print 1, but it was %q", output)
	}

	bf.Interpret("<<<<.")
	if output := bf.outputValue; output != "'\\x01'" {
		t.Errorf("Expected to print 1, but it was %q", output)
	}
}

func TestInterpretReadInput(t *testing.T) {
	t.Log("Test BF interpretter reading input")
	bf := NewInterpretter()
	bf.inputValue = 65
	bf.Interpret(",.")
	if output := bf.outputValue; output != "'A'" {
		t.Errorf("Expected to print A, but it was %q", output)
	}
}

func TestInterpretDebug(t *testing.T) {
	t.Log("Test BF Interpretter debugger")
	bf := NewInterpretter()
	pointer, tapeValue, tapeLength := bf.Debug()
	if pointer != 0 || tapeValue != 0 || tapeLength != 1 {
		t.Errorf("Interpretter values not correct pointer: %d cellValue: %d tapeLength: %d", pointer, tapeValue, tapeLength)
	}
	bf.Interpret("+++++ ++++")
	pointer, tapeValue, tapeLength = bf.Debug()
	if pointer != 0 || tapeValue != 9 || tapeLength != 1 {
		t.Errorf("Interpretter values not correct pointer: %d cellValue: %d tapeLength: %d", pointer, tapeValue, tapeLength)
	}

	bf.Interpret(">>+++++")
	pointer, tapeValue, tapeLength = bf.Debug()
	if pointer != 2 || tapeValue != 5 || tapeLength != 3 {
		t.Errorf("Interpretter values not correct pointer: %d cellValue: %d tapeLength: %d", pointer, tapeValue, tapeLength)
	}

	bf.Interpret("<")
	pointer, tapeValue, tapeLength = bf.Debug()
	if pointer != 1 || tapeValue != 0 || tapeLength != 3 {
		t.Errorf("Interpretter values not correct pointer: %d cellValue: %d tapeLength: %d", pointer, tapeValue, tapeLength)
	}
}

// Starting to implement loops
func TestInterpretLoop(t *testing.T) {
	t.Log("Test BF interpretter loop implementation")
	bf := NewInterpretter()

	bf.Interpret("+++++ ++++")
	pointer, tapeValue, tapeLength := bf.Debug()
	if pointer != 0 || tapeValue != 9 || tapeLength != 1 {
		t.Errorf("Interpretter values not correct pointer: %d cellValue: %d tapeLength: %d", pointer, tapeValue, tapeLength)
	}

	bf.Interpret("[ > +++++  +++++ < -] >.")
	pointer, tapeValue, tapeLength = bf.Debug()
	if pointer != 1 || tapeValue != 90 || tapeLength != 2 {
		t.Errorf("Interpretter values not correct pointer: %d cellValue: %d tapeLength: %d", pointer, tapeValue, tapeLength)
	}
}

func TestInterpretTwoLoops(t *testing.T) {
	t.Log("Test BF interpretter two loops")
	bf := NewInterpretter()
	bf.Interpret("+++[ > +++ [ > +++ < -] < -] >> .")
	pointer, tapeValue, tapeLength := bf.Debug()
	if pointer != 2 || tapeValue != 27 || tapeLength != 3 {
		t.Errorf("Interpretter values not correct pointer: %d cellValue: %d tapeLength: %d", pointer, tapeValue, tapeLength)
	}
}

func TestInterpretGetOutput(t *testing.T) {
	t.Log("Test BF interpretter get output")
	bf := NewInterpretter()
	bf.Interpret("+++++ +++++" +
		"[ > +++++ ++  > +++++ +++++   > +++  > + <<<< - ]" +
		"> ++ . > + . +++++ ++ . . +++ . > ++ . << +++++ +++++ +++++ ." +
		"> . +++ . ----- - . ----- --- . > + . > .")

	if output := bf.GetOutput(); output != "Hello World!\n" {
		t.Errorf("Expected to print A, but it was %q", output)
	}
}

func TestInterpretHelloWorld(t *testing.T) {
	t.Log("Test BF interpret hello world")
	bf := NewInterpretter()
	bf.Interpret("+++++ +++++" +
		"[ > +++++ ++  > +++++ +++++   > +++  > + <<<< - ]" +
		"> ++ . > + . +++++ ++ . . +++ . > ++ . << +++++ +++++ +++++ ." +
		"> . +++ . ----- - . ----- --- . > + . > .")

	if output := bf.GetOutput(); output != "Hello World!\n" {
		t.Errorf("Expected to print A, but it was %q", output)
	}

}
