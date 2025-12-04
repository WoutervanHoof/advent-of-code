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
		line := scanner.Text()

		tens := 0
		ones := 0
	inner:
		for i := 0; i < len(line)-1; i++ {
			digit, _ := strconv.Atoi(line[i : i+1])

			if tens < digit {
				tens = digit
				ones = 0
				continue inner
			}

			ones = max(ones, digit)
		}

		digit, _ := strconv.Atoi(line[len(line)-1:])
		ones = max(ones, digit)

		result += tens*10 + ones
		fmt.Printf("line: %d%d, result: %d\n", tens, ones, result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %d\n", result)
}
