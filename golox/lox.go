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

func (l *Lox) Run() error {
	for {
		line, err := l.reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
		l.source += line

		l.lines = append(l.lines, line)
	}

	scanner := NewScanner(l.source)

	tokens, _ := scanner.ScanTokens()

	fmt.Println(tokens)

	return nil
}
