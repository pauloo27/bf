package bf

import (
	"fmt"
)

/**
  a bad brainfuck interpreter.
*/

const (
	memorySize = 1024
)

type Program struct {
	Program      []byte
	ProgramIndex int
	MemoryIndex  int
	Memory       [memorySize]byte
	Input        []byte
	InputIndex   int
	OutputIndex  int
	LoopStart    int
}

func (p *Program) Run() (output []byte, err error) {
	programSize := len(p.Program)
	inputSize := len(p.Input)
	output = make([]byte, memorySize)
	for {
		if p.ProgramIndex >= programSize {
			break
		}
		switch p.Program[p.ProgramIndex] {
		case '>':
			p.MemoryIndex++
		case '<':
			p.MemoryIndex--
		case '+':
			p.Memory[p.MemoryIndex]++
		case '-':
			p.Memory[p.MemoryIndex]--
		case ',':
			if p.InputIndex >= inputSize {
				err = fmt.Errorf("input size is %d but tried to read at %d", inputSize, p.InputIndex)
				return
			}
			p.Memory[p.MemoryIndex] = p.Input[p.InputIndex]
			p.InputIndex++
		case '.':
			output[p.OutputIndex] = p.Memory[p.MemoryIndex]
			p.OutputIndex++
		case '[':
			p.LoopStart = p.ProgramIndex
		case ']':
			if p.Memory[p.MemoryIndex] != 0 {
				p.ProgramIndex = p.LoopStart
			}
		}
		p.ProgramIndex++
	}
	return
}

func NewProgram(code, input string) *Program {
	return &Program{
		Program:      []byte(code),
		ProgramIndex: 0,
		Input:        []byte(input),
	}
}
