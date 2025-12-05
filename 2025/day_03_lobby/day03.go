package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()

		digits := make([]byte, 12)
		copy(digits, line[len(line)-12:])

		// from back to front, check digits
		for i := len(line) - 13; i >= 0; i-- {
			digit := line[i]

			// swap until no longer necessary
			for j := 0; j < 12 && digit >= digits[j]; j++ {
				temp := digits[j]
				digits[j] = digit
				digit = temp
			}
		}

		partial, _ := strconv.Atoi(string(digits))
		result += partial
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %d\n", result)
}
