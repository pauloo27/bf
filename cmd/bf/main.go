package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Pauloo27/bf/pkg/bf"
)

var (
	filePath string
	input    string
)

func init() {
	flag.StringVar(&filePath, "file", "", "the program source code file path")
	flag.StringVar(&input, "input", "", "the program input")
}

func main() {
	flag.Parse()
	if filePath == "" {
		flag.Usage()
		os.Exit(1)
	}
	rawCode, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("cannot open source file")
		os.Exit(1)
	}
	pgm := bf.NewProgram(string(rawCode), input)
	out, err := pgm.Run()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	for i := 0; i < pgm.OutputIndex; i++ {
		fmt.Print(string(out[i]))
	}
}
