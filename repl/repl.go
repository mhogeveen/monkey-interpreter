package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey-interpreter/colorize"
	"monkey-interpreter/lexer"
	"monkey-interpreter/parser"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, colorize.Colorize(">> ", "green"))
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, colorize.Colorize("Woops! We ran into some monkey business here!\n", "blue"))
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, colorize.Colorize("- "+msg+"\n", "red"))
	}
}
