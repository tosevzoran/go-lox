package golox

import (
	"bufio"
	"fmt"
	"io"
)

type Lox struct {
	lines  []string
	source string
	reader *bufio.Reader
}

func New(r *bufio.Reader) *Lox {
	return &Lox{reader: r, lines: make([]string, 0), source: ""}
}

func (l *Lox) Run(interactive bool) error {
	for {
		line, err := l.reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if interactive {
			run(line)
			continue
		}

		l.source += line

		l.lines = append(l.lines, line)
	}

	if !interactive {
		run(l.source)
	}

	return nil
}

func run(source string) {
	scanner := NewScanner(source)

	tokens, err := scanner.ScanTokens()

	if err != nil {
		panic(fmt.Sprintf("error while scanning %v", err))
	}

	parser := NewParser(tokens)

	expressions, err := parser.parse()

	if err != nil {
		panic(fmt.Sprintf("error while parsing %v", err))
	}

	interpreter := NewInterpreter()

	val, err := interpreter.interpret(expressions)

	if err != nil {
		panic(fmt.Sprintf("error while interpreting %v", err))
	}

	fmt.Println(val)
}
