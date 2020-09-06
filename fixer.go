package fixer

import (
	"io"
	"sync"
)

// Fixer :
type Fixer struct {
	Writer     io.Writer
	PrefixFunc func(infix string) string
	InfixFunc  func(s string) string
	SuffixFunc func(infix string) string
}

var mux sync.Mutex

func (f Fixer) Write(p []byte) (n int, err error) {
	// TODO: error
	if f.Writer == nil {
		return len(p), err
	}
	go f.asyncWrite(p)
	return len(p), err
}

func (f Fixer) asyncWrite(p []byte) {
	mux.Lock()

	line := []byte{}

	afterNewline := false
	for _, b := range p {
		if afterNewline {
			io.WriteString(f.Writer, f.format(line))

			afterNewline = false
			line = []byte{}
		}
		if b == '\n' {
			afterNewline = true
		} else {
			line = append(line, b)
		}
	}
	io.WriteString(f.Writer, f.format(line))

	mux.Unlock()
}

func (f Fixer) format(line []byte) string {
	var infix = string(line)
	if f.InfixFunc != nil {
		infix = f.InfixFunc(string(line))
	}

	var prefix = ""
	if f.PrefixFunc != nil {
		prefix = f.PrefixFunc(infix)
	}

	var suffix = ""
	if f.SuffixFunc != nil {
		suffix = f.SuffixFunc(infix)
	}

	return prefix + infix + suffix + "\n"
}
