package bf

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
