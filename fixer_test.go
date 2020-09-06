package fixer

import (
	"testing"
)

var hello = []byte("hello")

func TestNoWriter(t *testing.T) {
	f := Fixer{}
	f.Write(hello)
}

func TestWriteReturns(t *testing.T) {
	f := Fixer{}
	n, err := f.Write(hello)
	if err != nil {
		t.Error("err returned")
	}
	if n != len(hello) {
		t.Errorf("n is not %d", len(hello))
	}
}

func TestWrite(t *testing.T) {
	var b = NewBufferSpy()

	f := Fixer{
		Writer: b,
	}
	f.Write(hello)
	<-b.writeChan

	expected := string(hello) + "\n"
	actual := string(b.buffer)

	if actual != expected {
		t.Errorf("actual: %s is not expected: %s", actual, expected)
	}
}

func TestCombos(t *testing.T) {
	tests := []struct {
		prefixFunc func(s string) string
		infixFunc  func(s string) string
		suffixFunc func(s string) string

		in  string
		out string
	}{
		{
			func(s string) string { return "prefix:" },
			nil,
			nil,
			"hello",
			"prefix:hello",
		},
		{
			nil,
			func(s string) string { return "infix:" + s + ":infix" },
			nil,
			"hello",
			"infix:hello:infix",
		},
		{
			nil,
			nil,
			func(s string) string { return ":suffix" },
			"hello",
			"hello:suffix",
		},
		{
			func(s string) string { return "prefix" },
			func(s string) string { return "infix" },
			func(s string) string { return "suffix" },
			"hello",
			"prefixinfixsuffix",
		},
	}

	for _, test := range tests {
		b := NewBufferSpy()

		f := Fixer{
			Writer:     b,
			PrefixFunc: test.prefixFunc,
			InfixFunc:  test.infixFunc,
			SuffixFunc: test.suffixFunc,
		}
		f.Write([]byte(test.in))
		<-b.writeChan

		expected := test.out + "\n"
		actual := string(b.buffer)

		if actual != expected {
			t.Errorf("actual: %s is not expected: %s", actual, expected)
		}
	}
}
