package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	ErrInvalidColumn = errors.New("invalid column number")
	ErrFileInvalid   = errors.New("cannot open file")
)

func sumOp(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}

	return sum
}

func avgOp(data []float64) float64 {
	if len(data) > 0 {
		avg := sumOp(data) / float64(len(data))
		return math.Round(avg*10000) / 10000
		//return avg
	} else {
		return 0.0
	}
}

type config struct {
	column int
	trim   bool
	avg    bool
	sum    bool
}

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintln(w, "Print column from STDIN or filename")
		fmt.Fprintln(w, "USAGE: coln 3 filename.txt")
		flag.PrintDefaults()
	}

	sum := flag.Bool("sum", false, "Calculate the sum of all numbers in the column")
	avg := flag.Bool("avg", false, "Calculate the average of all numbers in the column")
	trimQuotes := flag.Bool("q", false, "Strip quotes")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		exit(nil)
	}

	column, err := strconv.Atoi(args[0])
	if err != nil {
		exit(fmt.Errorf("%w: %q", ErrInvalidColumn, args[0]))
	}

	cfg := config{
		column: column,
		trim:   *trimQuotes,
		sum:    *sum,
		avg:    *avg,
	}

	var r io.Reader

	if len(args) == 1 {
		// PIPE
		r = os.Stdin
	} else {
		// FILE
		r, err = os.Open(args[1])
		if err != nil {
			exit(fmt.Errorf("%s: %w", ErrFileInvalid.Error(), err))
		}
	}

	run(r, os.Stdout, cfg)
}

func run(r io.Reader, w io.Writer, cfg config) {
	var lines = bufio.NewScanner(r)
	lines.Split(bufio.ScanLines)
	var data []float64

	computeStats := cfg.avg || cfg.sum

	for lines.Scan() {
		words := bufio.NewScanner(bytes.NewReader(lines.Bytes()))
		words.Split(bufio.ScanWords)
		c := 0
		var word string
		for words.Scan() {
			c++
			if cfg.column == c {
				word = words.Text()
				if cfg.trim {
					word = strings.Trim(word, "\";'")
				}

				if computeStats {
					num, err := strconv.ParseFloat(word, 64)
					if err != nil {
						// silently skip invalid numbers
						continue
					}
					data = append(data, num)
				} else {
					fmt.Fprintln(w, word)
				}
			}
		}
	}

	if len(data) > 0 {
		if cfg.sum {
			fmt.Fprintln(w, sumOp(data))
		} else if cfg.avg {
			fmt.Fprintln(w, avgOp(data))
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
