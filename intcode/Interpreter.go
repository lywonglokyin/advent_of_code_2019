package intcode

import (
	"fmt"
)

// A Interpreter handles Intcode
type Interpreter struct {
	intcodes []int64
	pc       uint64 //program counter
}

// NewIntcodeInterpreter is constructor of Interpreter
func NewIntcodeInterpreter(intcodes []int64) *Interpreter {
	interpreter := new(Interpreter)
	interpreter.intcodes = intcodes
	interpreter.pc = 0
	return interpreter
}

// Reset resets the interpreter
func (interpreter *Interpreter) Reset() {
	interpreter.pc = 0
}

// Execute execute the Intcode starting from current pc.
func (interpreter *Interpreter) Execute() {
	for {
		pos := interpreter.pc
		opcode := interpreter.intcodes[pos] % 100
		if opcode == 99 {
			interpreter.pc++
			break
		}

		switch opcode {
		case 1, 2, 7, 8:
			interpreter.pc += 4
			interpreter.executeCommand(opcode, interpreter.intcodes[pos], interpreter.intcodes[pos+1], interpreter.intcodes[pos+2], interpreter.intcodes[pos+3])
		case 3, 4:
			interpreter.pc += 2
			interpreter.executeShortCommand(opcode, interpreter.intcodes[pos], interpreter.intcodes[pos+1])
		case 5, 6:
			interpreter.pc += 3
			interpreter.execute3Args(opcode, interpreter.intcodes[pos], interpreter.intcodes[pos+1], interpreter.intcodes[pos+2])
		default:
			panic("Unknown opcode!")
		}
	}
}

// At return the intcode at pos
func (interpreter *Interpreter) At(pos uint64) int64 {
	return interpreter.intcodes[pos]
}

// Intcodes is a getter for intcodes
func (interpreter *Interpreter) Intcodes() []int64 {
	return interpreter.intcodes
}

func (interpreter *Interpreter) executeCommand(opcode int64, modeCode int64, arg1 int64, arg2 int64, arg3 int64) {
	mode1 := (modeCode % 1000) / 100
	mode2 := (modeCode % 10000) / 1000

	var val1, val2 int64
	val1 = interpreter.getValue(arg1, mode1)
	val2 = interpreter.getValue(arg2, mode2)

	switch opcode {
	case 1:
		additionResult := val1 + val2
		interpreter.intcodes[arg3] = additionResult
	case 2:
		multiplicationResult := val1 * val2
		interpreter.intcodes[arg3] = multiplicationResult
	case 7:
		if val1 < val2 {
			interpreter.intcodes[arg3] = 1
		} else {
			interpreter.intcodes[arg3] = 0
		}
	case 8:
		if val1 == val2 {
			interpreter.intcodes[arg3] = 1
		} else {
			interpreter.intcodes[arg3] = 0
		}
	default:
		panic("Unknown opcode!")
	}
}

func (interpreter *Interpreter) executeShortCommand(opcode int64, modeCode int64, arg1 int64) {
	switch opcode {
	case 3:
		var i int64
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			panic(err)
		}
		interpreter.intcodes[arg1] = i
	case 4:
		mode1 := (modeCode % 1000) / 100
		val1 := interpreter.getValue(arg1, mode1)
		fmt.Println(val1)
	default:
		panic("Unknown opcode!")
	}
}

func (interpreter *Interpreter) execute3Args(opcode int64, modeCode int64, arg1 int64, arg2 int64) {
	mode1 := (modeCode % 1000) / 100
	mode2 := (modeCode % 10000) / 1000

	var val1, val2 int64
	val1 = interpreter.getValue(arg1, mode1)
	val2 = interpreter.getValue(arg2, mode2)
	switch opcode {
	case 5:
		if val1 != 0 {
			interpreter.pc = uint64(val2)
		}
	case 6:
		if val1 == 0 {
			interpreter.pc = uint64(val2)
		}
	default:
		panic("Unknown opcode!")
	}
}

func (interpreter *Interpreter) getValue(intcode int64, mode int64) int64 {
	switch mode {
	case 0:
		return interpreter.intcodes[intcode]
	case 1:
		return intcode
	default:
		panic("Invalid mode!")
	}
	return -1
}
