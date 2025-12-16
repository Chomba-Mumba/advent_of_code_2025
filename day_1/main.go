package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	//create scanner to read file line by line
	scanner := bufio.NewScanner(file)

	dial := 50
	out := 0

	// read each line in file
	for scanner.Scan() {
		line := scanner.Text()
		dial += parse_line(line)

		dial = (dial + 100) % 100

		if dial == 0 {
			out += 1
		}
		fmt.Printf("dial: %v - out: %v - line: %v\n", dial, out, line)
	}

	//check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	fmt.Printf("final output %v\n", out)
}

func parse_line(line string) int {
	out, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatalf("unable to convert string to int")
	}
	if line[0] == 'R' {
		return out
	} else {
		return -out
	}
}
