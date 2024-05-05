package golox

import (
	"bufio"
	"io"
)

type Lox struct {
	lines  []string
	reader *bufio.Reader
}

func New(r *bufio.Reader) *Lox {
	return &Lox{reader: r, lines: make([]string, 0)}
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

		l.lines = append(l.lines, line)
	}

	return nil
}
