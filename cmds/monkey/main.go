package main


import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"os"
)

func runFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Could not open file %s: %s\n", fileName, err)
		os.Exit(1)
	}
	defer file.Close()

	env := object.NewEnvironment()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(os.Stdout, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %s\n", fileName, err)
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: monkey <file>")
		os.Exit(1)
	}
	fileName := os.Args[1]
	runFile(fileName)
}
