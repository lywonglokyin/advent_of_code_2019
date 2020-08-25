package intcode

import "log"

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
func (interpreter Interpreter) Reset() {
	interpreter.pc = 0
}

// Execute execute the Intcode starting from current pc.
func (interpreter Interpreter) Execute() {
	for {
		pos := interpreter.pc
		opcode := interpreter.intcodes[pos]
		if opcode == 99 {
			interpreter.pc++
			break
		}
		interpreter.pc += 4
		interpreter.executeCommand(interpreter.intcodes[pos], interpreter.intcodes[pos+1], interpreter.intcodes[pos+2], interpreter.intcodes[pos+3])
	}
}

// At return the intcode at pos
func (interpreter Interpreter) At(pos uint64) int64 {
	return interpreter.intcodes[pos]
}

func (interpreter Interpreter) Intcodes() []int64 {
	return interpreter.intcodes
}

func (interpreter Interpreter) executeCommand(opcode int64, arg1 int64, arg2 int64, arg3 int64) {
	switch opcode {
	case 1:
		additionResult := interpreter.intcodes[arg1] + interpreter.intcodes[arg2]
		interpreter.intcodes[arg3] = additionResult
	case 2:
		multiplicationResult := interpreter.intcodes[arg1] * interpreter.intcodes[arg2]
		interpreter.intcodes[arg3] = multiplicationResult
	default:
		log.Fatal("Unknown opcode!")
	}
}
