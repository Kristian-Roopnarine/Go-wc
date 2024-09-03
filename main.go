package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var showBytes = false
var showLines = false
var showWords = false

func main() {
	filename := parseArgs(os.Args[1:])

	var file io.Reader
	var err error

	if filename == "" {
		file = bufio.NewReader(os.Stdin)
	} else {
		file, err = os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(filename)

	input := ""
	buf := make([]byte, 4096)
	var totalBytes = 0
	for {
		n, err := file.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Printf("End of file\n")
				break
			}
			log.Fatal(err)
			return
		}

		totalBytes += n

	}

	words := readWords(input)
	lines := readLines(input)

	if showBytes {
		fmt.Println(totalBytes)
	}

	if showLines {
		fmt.Println(lines)
	}

	if showWords {
		fmt.Println(words)
	}

}

func parseArgs(args []string) string {
	var inflag = false

	for idx, arg := range args {
		if idx == len(args)-1 {
			return args[idx]
		}

		for _, c := range arg {
			if '-' == c {
				inflag = true
				continue
			}

			if inflag {
				switch c {
				case 'c':
					showBytes = true
				case 'l':
					showLines = true
				case 'w':
					showWords = true
				}
			}
		}

		inflag = false
	}

	return ""

}

func readWords(s string) int {
	return len(strings.Fields(s))
}

func readLines(s string) int {
	count := 0
	for _, c := range s {
		if c == '\n' {
			count++
		}
	}
	return count
}
