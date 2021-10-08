// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// stdin, filenames
// map of maps
// inputName -> line -> count

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		counts["stdin"] = make(map[string]int)
		countLines(os.Stdin, counts["stdin"])
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			counts[arg] = make(map[string]int)
			countLines(f, counts[arg])
			f.Close()
		}
	}
	for inputSource, counts := range counts {
		var hasDuplicates bool
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
				hasDuplicates = true
			}
		}
		if hasDuplicates && inputSource != "stdin" {
			fmt.Printf("%s has duplicate lines.\n", inputSource)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
