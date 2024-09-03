package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

	buf := make([]byte, 4096)
	var totalBytes = 0
	input := ""
	for {
		n, err := file.Read(buf)
		totalBytes += n
		input += string(buf[:n])

		if err != nil {
			if err == io.EOF {
				fmt.Printf("End of file\n")
				break
			}
			log.Fatal(err)
			return
		}

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

		if idx == len(args)-1 {
			if inflag {
				break
			}
			return args[idx]
		}

		inflag = false
	}
	return ""

}

func readWords(s string) int {
	count := 0
	var inword = false
	for _, c := range s {
		switch c {
		case '\r', '\t', ' ', '\n':
			if inword {
				count++
				inword = false
			}
		default:
			inword = true
		}

	}
	return count
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
