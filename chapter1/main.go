package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "count lines")
	bts := flag.Bool("b", false, "count number of bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bts))
}

func count(r io.Reader, countLine bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLine && !countBytes {
		scanner.Split(bufio.ScanWords)
	}
	wc := 0

	for scanner.Scan() {
		if countBytes {
			wc += len(scanner.Bytes())
		} else {
			wc++
		}
	}

	return wc
}

// to run = echo "My first command line program" | ./chapter1
// to count line use printf instead of echo = printf "line1 \nline2 \n" | ./chapter1 -l
