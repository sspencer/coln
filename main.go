package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintln(w, "Print column from STDIN or filename")
		fmt.Fprintln(w, "USAGE: coln 3 filename.txt")
		flag.PrintDefaults()
	}

	trimQuotes := flag.Bool("q", false, "Strip quotes")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		exit(nil)
	}

	column, err := strconv.Atoi(args[0])
	if err != nil {
		exit(fmt.Errorf("please specify column as an integer"))
	}

	var r io.Reader

	if len(args) == 1 {
		// PIPE
		r = os.Stdin
	} else {
		// FILE
		r, err = os.Open(args[1])
		if err != nil {
			exit(err)
		}
	}

	run(r, os.Stdout, column, *trimQuotes)
}

func run(r io.Reader, w io.Writer, column int, trimQuotes bool) {
	var lines = bufio.NewScanner(r)
	lines.Split(bufio.ScanLines)
	for lines.Scan() {
		words := bufio.NewScanner(bytes.NewReader(lines.Bytes()))
		words.Split(bufio.ScanWords)
		c := 0
		for words.Scan() {
			c++
			if column == c {
				if trimQuotes {
					fmt.Fprintln(w, strings.Trim(words.Text(), "\";'"))
				} else {
					fmt.Fprintln(w, words.Text())
				}
			}
		}
	}
}

func exit(err error) {
	if err == nil {
		flag.Usage()
	} else {
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(1)
}
