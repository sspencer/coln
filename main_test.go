package main

import (
	"bytes"
	"strings"
	"testing"
)

const (
	noTrimInput = `
1 2 3
one two three`
	noTrimOutput = `2
two
`
	trimInput = `v1 = "abc123";
v2 = "cafe456";`
	trimOutput = `abc123
cafe456
`
)

func TestNoTrim(t *testing.T) {
	input := strings.NewReader(noTrimInput)
	var output bytes.Buffer

	run(input, &output, 2, true)
	if noTrimOutput != output.String() {
		t.Fatalf("Result %q does not match %q", output.String(), noTrimOutput)
	}
}

func TestTrim(t *testing.T) {
	input := strings.NewReader(trimInput)
	var output bytes.Buffer

	run(input, &output, 3, true)
	if trimOutput != output.String() {
		t.Fatalf("Result %q does not match %q", output.String(), trimOutput)
	}
}
